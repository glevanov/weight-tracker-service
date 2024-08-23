import { createServer } from "node:http";
import { Router } from "node-router";

import { config } from "./config.js";
import { getWeights } from "./handlers/get-weights.js";
import { addWeight } from "./handlers/add-weight.js";
import { login } from "./handlers/login.js";

const router = new Router();

router.addRoute("GET", "/health-check", (req, res) => {
  res.writeHead(200, { "Content-Type": "text/plain" });
  res.end("OK");
});

router.addRoute("GET", "/weights", (req, res) => {
  void getWeights(req.url, (result) => {
    if (result.isSuccess) {
      res.writeHead(200, { "Content-Type": "application/json" });
      res.end(JSON.stringify(result.data));
    } else {
      res.writeHead(500, { "Content-Type": "text/plain" });
      res.end(result.error.toString());
    }
  });
});

router.addRoute("POST", "/weights", (req, res) => {
  let body = "";
  req.on("data", (chunk) => {
    body += chunk.toString();
  });

  req.on("end", () => {
    void addWeight(body, (result) => {
      if (result.isSuccess) {
        res.writeHead(201, { "Content-Type": "text/plain" });
        res.end(JSON.stringify(result.data));
      } else {
        res.writeHead(500, { "Content-Type": "text/plain" });
        res.end(result.error.toString());
      }
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
        res.writeHead(200, { "Content-Type": "text/plain" });
        res.end(JSON.stringify(result.data));
      } else {
        res.writeHead(500, { "Content-Type": "text/plain" });
        res.end(result.error.toString());
      }
    });
  });
});

/* Registration is closed :)
router.addRoute("POST", "/register", (req, res) => {
  let body = "";
  req.on("data", (chunk) => {
    body += chunk.toString();
  });

  req.on("end", () => {
    registerUser(body, (result) => {
      if (result.isSuccess) {
        res.writeHead(201, { "Content-Type": "text/plain" });
        res.end(JSON.stringify(result.data));
      } else {
        res.writeHead(500, { "Content-Type": "text/plain" });
        res.end(result.error.toString());
      }
    });
  });
});*/

const server = createServer((req, res) => {
  res.setHeader("Access-Control-Allow-Origin", config.frontendUrl);
  router.handle(req, res);
});

server.listen(config.port, () => {
  console.log(`Server is running on http://localhost:${config.port}`);
});
