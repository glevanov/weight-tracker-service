import { Connection } from "../connection.js";
import type { Result, Token } from "../types.js";
import {
  validateAndFormatWeight,
  WeightValidationError,
} from "../validation/validation.js";
import { Lang, locales } from "../i18n/i18n.js";

export const addWeight = async (
  body: string,
  token: Token,
  lang: Lang,
  callback: (result: Result<string>) => void,
) => {
  const connection = Connection.instance;

  let parsedBody: Record<string, unknown>;

  try {
    parsedBody = JSON.parse(body);
  } catch {
    callback({
      isSuccess: false,
      error: locales[lang].validation.weight.failedToParse,
    });
    return;
  }

  const validationResult = validateAndFormatWeight(
    String(parsedBody.weight),
    lang,
  );

  if (validationResult instanceof WeightValidationError) {
    callback({
      isSuccess: false,
      error: validationResult.message,
    });
    return;
  }

  const result = await connection.addWeight(
    validationResult,
    token.username,
    lang,
  );

  callback(result);
};
