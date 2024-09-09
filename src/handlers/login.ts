import { Connection } from "../connection.js";
import { Result } from "../types.js";
import { Lang, locales } from "../i18n/i18n.js";

export const login = async (
  body: string,
  lang: Lang,
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
      error: locales[lang].validation.auth.failedToParse,
    });
    return;
  }

  const result = await connection.loginUser(username, password, lang);

  callback(result);
};
