import { createServer } from "node:http";
import { Router } from "node-router";

import { config } from "./config.js";
import { getWeights } from "./get-weights.js";
import { addWeight } from "./add-weight.js";

const router = new Router();

router.addRoute("GET", "/health-check", (req, res) => {
  res.writeHead(200, { "Content-Type": "text/plain" });
  res.end("OK");
});

router.addRoute("GET", "/weights", (req, res) => {
  getWeights((result) => {
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
    addWeight(body, (result) => {
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

const server = createServer((req, res) => {
  res.setHeader("Access-Control-Allow-Origin", config.frontendUrl);
  router.handle(req, res);
});

server.listen(config.port, () => {
  console.log(`Server is running on http://localhost:${config.port}`);
});
