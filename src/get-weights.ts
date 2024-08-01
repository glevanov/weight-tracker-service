import { Connection } from "./connection.js";
import type { Result, Weight } from "./types.js";

export const getWeights = async (
  callback: (result: Result<Weight[]>) => void,
) => {
  const connection = Connection.instance;

  const result = await connection.getWeights();

  callback(result);
};
