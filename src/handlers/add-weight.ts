import { Connection } from "../connection.js";
import type { Result } from "../types.js";
import { validateAndFormatWeight, WeightValidationError } from "../validate.js";
import { literals } from "../literals.js";

export const addWeight = async (
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
      error: new Error(literals.validation.weight.failedToParse),
    });
    return;
  }

  const validationResult = validateAndFormatWeight(String(parsedBody.weight));

  if (validationResult instanceof WeightValidationError) {
    callback({
      isSuccess: false,
      error: validationResult,
    });
    return;
  }

  const result = await connection.addWeight(validationResult);

  callback(result);
};
