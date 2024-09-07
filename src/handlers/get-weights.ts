import type { IncomingMessage } from "http";

import { Connection } from "../connection.js";
import type { Result, Weight } from "../types.js";
import {
  validateAndParseTimestamp,
  TimestampValidationError,
} from "../validation/validation.js";
import { literals } from "../literals.js";

export const getWeights = async (
  url: IncomingMessage["url"],
  callback: (result: Result<Weight[]>) => void,
) => {
  const connection = Connection.instance;

  const params = new URLSearchParams(url?.split("?")[1]);

  const start = validateAndParseTimestamp(params.get("start"));
  const end = validateAndParseTimestamp(params.get("end"));

  const errors: string[] = [];

  if (start instanceof TimestampValidationError) {
    errors.push(
      `${literals.validation.timestamp.failedToParseStart}: ${start.message}`,
    );
  }
  if (end instanceof TimestampValidationError) {
    errors.push(
      `${literals.validation.timestamp.failedToParseEnd}: ${end.message}`,
    );
  }

  if (errors.length > 0) {
    callback({
      isSuccess: false,
      error: errors.join("; "),
    });
    return;
  }

  const result = await connection.getWeights(start as Date, end as Date);

  callback(result);
};
