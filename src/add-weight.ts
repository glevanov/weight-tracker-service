import { Connection } from "./connection.js";
import type { Result } from "./types.js";

export const addWeight = async (
  body: string,
  callback: (result: Result<string>) => void,
) => {
  const connection = Connection.instance;

  const parsedBody = JSON.parse(body);
  const weight = Number(parsedBody.weight);

  if (isNaN(weight)) {
    callback({
      isSuccess: false,
      error: new Error(
        `Incorrect weight type. Expected: number. Actual: ${weight}`,
      ),
    });
  }

  const result = await connection.addWeight(weight);

  callback(result);
};
