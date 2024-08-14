import type { Result } from "../types.js";
import { Connection } from "../connection.js";
import { literals } from "../literals.js";

export const registerUser = async (
  body: string,
  callback: (result: Result<string>) => void,
) => {
  const connection = Connection.instance;

  let parsedBody: Record<string, unknown>;

  try {
    parsedBody = JSON.parse(body);
  } catch {
    callback({
      isSuccess: false,
      error: new Error(literals.validation.auth.failedToParse),
    });
    return;
  }

  const { username, password } = parsedBody;
  if (typeof username !== "string" || typeof password !== "string") {
    callback({
      isSuccess: false,
      error: new Error(literals.validation.auth.invalidFormat),
    });
    return;
  }

  const result = await connection.registerUser(username, password);

  callback(result);
};
