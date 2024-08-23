import { randomBytes, scrypt } from "node:crypto";
import type { IncomingMessage, ServerResponse } from "node:http";

import { literals } from "../literals.js";

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
    const authHeader = req.headers.authorization;
    if (typeof authHeader !== "string") {
      throw new Error();
    }
    const token = authHeader.replace("Bearer ", "");
    if (!token) {
      throw new Error();
    }
    // to be implemented
  } catch {
    res.writeHead(401, { "Content-Type": "text/plain" });
    res.end(literals.error.user.failedToAuthorize);
  }
};

export const parseCookies = (req: IncomingMessage) => {
  const list: Record<string, string> = {};
  const cookieHeader = req.headers.cookie;

  if (cookieHeader) {
    try {
      const cookies = cookieHeader.split(";");
      for (const cookie of cookies) {
        // eslint-disable-next-line prefer-const
        let [name, ...rest] = cookie.split("=");
        name = name.trim();
        if (!name) return;
        const value = rest.join("=").trim();
        if (!value) return;
        list[name] = decodeURIComponent(value);
      }
    } catch {
      // do nothing
    }
  }

  return list;
};
