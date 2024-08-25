import type { IncomingMessage, ServerResponse } from "node:http";
import { Connection } from "../connection.js";
import { literals } from "../literals.js";

const parseCookieHeader = (req: IncomingMessage) => {
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
export const authMiddleware = async (
  req: IncomingMessage,
  res: ServerResponse,
) => {
  try {
    const { session } = parseCookieHeader(req);
    if (typeof session !== "string") {
      throw new Error();
    }
    const connection = Connection.instance;
    const hasSession = await connection.hasSession(session);

    if (!hasSession) {
      throw new Error();
    }
  } catch {
    res.writeHead(401, { "Content-Type": "text/plain" });
    res.end(literals.error.user.failedToAuthorize);
  }
};
