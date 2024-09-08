import {
  createServer,
  type IncomingMessage,
  type ServerResponse,
} from "node:http";
import { Router } from "node-router";

import { config } from "./config.js";
import { getWeights } from "./handlers/get-weights.js";
import { addWeight } from "./handlers/add-weight.js";
import { login } from "./handlers/login.js";
import { getCookieHeader } from "./auth/auth.js";
import { authMiddleware } from "./auth/authMiddleware.js";
import { literals } from "./literals.js";
import { SuccessResult } from "./types.js";

const OK: SuccessResult<string> = {
  isSuccess: true,
  data: "OK",
};

const router = new Router();

router.addRoute("GET", "/health-check", (req, res) => {
  res.writeHead(200, { "Content-Type": "application/json" });
  res.end(JSON.stringify(OK));
});

router.addRoute("GET", "/weights", (req, res) => {
  authMiddleware(req, res).then(() => {
    if (res.writableEnded) return;

    void getWeights(req.url, (result) => {
      if (result.isSuccess) {
        res.writeHead(200, { "Content-Type": "application/json" });
      } else {
        res.writeHead(400, { "Content-Type": "application/json" });
      }
      res.end(JSON.stringify(result));
    });
  });
});

router.addRoute("OPTIONS", "/weights", (req, res) => {
  res.writeHead(204, {
    "Access-Control-Allow-Methods": "POST, OPTIONS",
    "Access-Control-Allow-Headers": "Content-Type",
  });
  res.end();
});

router.addRoute("POST", "/weights", async (req, res) => {
  authMiddleware(req, res).then(() => {
    if (res.writableEnded) return;

    let body = "";
    req.on("data", (chunk) => {
      body += chunk.toString();
    });

    req.on("end", () => {
      void addWeight(body, (result) => {
        if (result.isSuccess) {
          res.writeHead(201, { "Content-Type": "application/json" });
        } else {
          res.writeHead(400, { "Content-Type": "application/json" });
        }
        res.end(JSON.stringify(result));
      });
    });
  });
});

router.addRoute("OPTIONS", "/login", (req, res) => {
  res.writeHead(204, {
    "Access-Control-Allow-Methods": "POST, OPTIONS",
    "Access-Control-Allow-Headers": "Content-Type",
  });
  res.end();
});

router.addRoute("POST", "/login", (req, res) => {
  let body = "";
  req.on("data", (chunk) => {
    body += chunk.toString();
  });

  req.on("end", () => {
    void login(body, (result) => {
      if (result.isSuccess) {
        res.writeHead(200, {
          "Content-Type": "application/json",
          "Set-Cookie": getCookieHeader(result.data),
          SameSite: "None",
        });
        res.end(JSON.stringify(OK));
      } else {
        res.writeHead(400, { "Content-Type": "application/json" });
        res.end(JSON.stringify(result));
      }
    });
  });
});

router.addRoute("GET", "/session-check", (req, res) => {
  authMiddleware(req, res).then(() => {
    if (res.writableEnded) return;

    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify(OK));
  });
});

const server = createServer((req: IncomingMessage, res: ServerResponse) => {
  res.setHeader("Access-Control-Allow-Origin", config.frontendUrl);
  res.setHeader("Access-Control-Allow-Credentials", "true");
  try {
    router.handle(req, res);
  } catch {
    res.writeHead(500, { "Content-Type": "application/json" });
    res.end({ isSuccess: false, error: literals.error.unknown });
  }
});

server.listen(config.port, () => {
  console.log(`Server is running on http://localhost:${config.port}`);
});
