import { createServer } from "node:http";

import { config } from "./config.js";
import { getWeights } from "./get-weights.js";
import { addWeight } from "./add-weight.js";

const server = createServer((req, res) => {
  res.setHeader("Access-Control-Allow-Origin", config.frontendUrl);

  if (req.url === "/health-check" && req.method === "GET") {
    res.writeHead(200, { "Content-Type": "text/plain" });
    res.end("OK");
  } else if (req.url === "/weights" && req.method === "GET") {
    getWeights((result) => {
      if (result.isSuccess) {
        res.writeHead(200, { "Content-Type": "application/json" });
        res.end(JSON.stringify(result.data));
      } else {
        res.writeHead(500, { "Content-Type": "text/plain" });
        res.end(result.error.toString());
      }
    });
  } else if (req.url === "/weights" && req.method === "POST") {
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
  } else {
    res.writeHead(404, { "Content-Type": "text/plain" });
    res.end("Not found");
  }
});

server.listen(config.port, () => {
  console.log(`Server is running on http://localhost:${config.port}`);
});
