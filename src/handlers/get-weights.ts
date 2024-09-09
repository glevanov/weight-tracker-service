import type { IncomingMessage } from "http";

import { Connection } from "../connection.js";
import type { Result, Token, Weight } from "../types.js";
import {
  validateAndParseTimestamp,
  TimestampValidationError,
} from "../validation/validation.js";
import { Lang, locales } from "../i18n/i18n.js";

export const getWeights = async (
  url: IncomingMessage["url"],
  token: Token,
  lang: Lang,
  callback: (result: Result<Weight[]>) => void,
) => {
  const connection = Connection.instance;

  const params = new URLSearchParams(url?.split("?")[1]);

  const start = validateAndParseTimestamp(params.get("start"), lang);
  const end = validateAndParseTimestamp(params.get("end"), lang);

  const errors: string[] = [];

  if (start instanceof TimestampValidationError) {
    errors.push(
      `${locales[lang].validation.timestamp.failedToParseStart}: ${start.message}`,
    );
  }
  if (end instanceof TimestampValidationError) {
    errors.push(
      `${locales[lang].validation.timestamp.failedToParseEnd}: ${end.message}`,
    );
  }

  if (errors.length > 0) {
    callback({
      isSuccess: false,
      error: errors.join("; "),
    });
    return;
  }

  const result = await connection.getWeights(
    start as Date,
    end as Date,
    token.username,
    lang,
  );

  callback(result);
};
