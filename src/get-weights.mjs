import { Connection } from "./connection.mjs";

/**
 * Get all weights from the database
 * @param {(result: { isSuccess: true, data: { date: string, weight: number }[] } | { isSuccess: false, error: Error }) => void} callback
 * @returns {Promise<{ isSuccess: true, data: { date: string, weight: number }[] } | { isSuccess: false, error: Error }>}
 */
export const getWeights = async (callback) => {
  const connection = Connection.instance;

  const result = await connection.getWeights();

  callback(result);
};
