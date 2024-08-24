import { randomBytes, scrypt } from "node:crypto";
import type { IncomingMessage, ServerResponse } from "node:http";

import { literals } from "../literals.js";
import { config } from "../config.js";

export const hashPassword = (
  password: string,
  salt: string,
): Promise<string | Error> => {
  return new Promise((resolve, reject) => {
    scrypt(password, salt, 64, (err, derivedKey) => {
      if (err) reject(err);
      resolve(derivedKey.toString("hex"));
    });
  });
};

export const generateSalt = () => randomBytes(16).toString("hex");

export const generateSessionId = () => randomBytes(16).toString("hex");

export const handleAuthorization = (
  req: IncomingMessage,
  res: ServerResponse,
) => {
  try {
    const { session } = parseCookieHeader(req);
    if (typeof session !== "string") {
      throw new Error();
    }
    // to be implemented
  } catch {
    res.writeHead(401, { "Content-Type": "text/plain" });
    res.end(literals.error.user.failedToAuthorize);
  }
};

export const getCookieHeader = (sessionId: string) => {
  return `session=${encodeURIComponent(sessionId)}; HttpOnly; Max-Age=${config.sessionDuration}`;
};

export const parseCookieHeader = (req: IncomingMessage) => {
  const list: Record<string, string> = {};

  const header = req.headers.cookie;

  if (typeof header === "string") {
    try {
      const cookies = header.split(";");
      for (const cookie of cookies) {
        // eslint-disable-next-line prefer-const
        let [name, ...rest] = cookie.split("=");
        name = name.trim();
        if (!name) continue;
        const value = rest.join("=").trim();
        if (!value) continue;
        list[name] = decodeURIComponent(value);
      }
    } catch {
      // do nothing
    }
  }

  return list;
};
