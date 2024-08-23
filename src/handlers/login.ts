import { Connection } from "../connection.js";
import { Result } from "../types.js";
import { literals } from "../literals.js";

export const login = async (
  body: string,
  callback: (result: Result<string>) => void,
) => {
  const connection = Connection.instance;

  let username: string;
  let password: string;

  try {
    const parsedBody = JSON.parse(body);
    const parsedUsername = parsedBody.username;
    const parsedPassword = parsedBody.password;

    if (
      typeof parsedUsername !== "string" ||
      typeof parsedPassword !== "string"
    ) {
      throw new Error();
    }
    username = parsedUsername;
    password = parsedPassword;
  } catch {
    callback({
      isSuccess: false,
      error: new Error(literals.validation.auth.failedToParse),
    });
    return;
  }

  const result = await connection.loginUser(username, password);

  callback(result);
};
