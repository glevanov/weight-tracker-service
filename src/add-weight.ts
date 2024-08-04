import { Connection } from "./connection.js";
import type { Result } from "./types.js";
import { validateAndFormatWeight, WeightValidationError } from "./validate.js";

export const addWeight = async (
  body: string,
  callback: (result: Result<string>) => void,
) => {
  const connection = Connection.instance;

  const parsedBody = JSON.parse(body);
  const validationResult = validateAndFormatWeight(parsedBody.weight);

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
