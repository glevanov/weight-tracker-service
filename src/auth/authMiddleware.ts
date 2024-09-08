import type { IncomingMessage, ServerResponse } from "node:http";
import jwt from "jsonwebtoken";
import { literals } from "../literals.js";
import { config } from "../config.js";
import { Token } from "../types.js";

const parseAuthHeader = (req: IncomingMessage) => {
  const authHeader = req.headers.authorization;
  if (typeof authHeader !== "string") {
    return null;
  }

  const [type, token] = authHeader.split(" ");
  if (type !== "Bearer" || typeof token !== "string" || token === "null") {
    return null;
  }

  return token;
};

const validateToken = (token: unknown): token is Token => {
  if (typeof token !== "object" || token === null) {
    return false;
  }

  if (!("username" in token) || !("iat" in token) || !("exp" in token)) {
    return false;
  }

  if (
    typeof token.username !== "string" ||
    typeof token.iat !== "number" ||
    typeof token.exp !== "number"
  ) {
    return false;
  }

  return true;
};

export const authMiddleware = async (
  req: IncomingMessage,
  res: ServerResponse,
) => {
  try {
    const token = parseAuthHeader(req);
    if (typeof token !== "string") {
      throw new Error();
    }

    const decoded = jwt.verify(token, config.jwtSecret);

    if (!validateToken(decoded)) {
      throw new Error();
    }

    return decoded;
  } catch {
    res.writeHead(401, { "Content-Type": "application/json" });
    res.end(
      JSON.stringify({
        isSuccess: false,
        error: literals.error.user.failedToAuthorize,
      }),
    );
  }
};
