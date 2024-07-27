import { getWeights } from "./src/get-weight-points.mjs";
import { createServer } from "node:http";
import { config } from "./config.mjs";

const server = createServer((req, res) => {
  res.setHeader("Access-Control-Allow-Origin", config.frontendUrl);

  let body = "";
  if (req.method === "POST") {
    req.on("data", (chunk) => {
      body += chunk.toString();
    });
  }

  req.on("end", () => {
    if (req.url === "/weights" && req.method === "GET") {
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
      const parsedBody = JSON.parse(body);
      const weight = Number(parsedBody.weight);

      console.log(weight);

      res.writeHead(200, { "Content-Type": "text/plain" });
      res.end(String(weight));
    } else {
      res.writeHead(404, { "Content-Type": "text/plain" });
      res.end("Not found");
    }
  });
});

server.listen(config.port, () => {
  console.log(`Server is running on http://localhost:${config.port}`);
});
