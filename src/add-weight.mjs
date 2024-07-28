import { Connection } from "./connection.mjs";

export const addWeight = async (body, callback) => {
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
