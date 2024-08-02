import { loadEnvFile } from "node:process";

loadEnvFile();

if (typeof process.env.CONNECTION_URI !== "string") {
  throw new Error("Connection URI is not provided");
}

export const config = {
  port: Number(process.env.PORT) || 3000,
  frontendUrl: process.env.FRONTEND_URL || "*",
  connectionUri: process.env.CONNECTION_URI,
  dbName: "weight-tracker",
  collectionName: "weight",
};
