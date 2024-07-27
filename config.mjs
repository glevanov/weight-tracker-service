import { loadEnvFile } from "node:process";

loadEnvFile();

export const config = {
  port: process.env.PORT || 3000,
  frontendUrl: process.env.FRONTEND_URL || "*",
  connectionUri: process.env.CONNECTION_URI,
  dbName: "weight-tracker",
  collectionName: "weight",
};

if (typeof config.connectionUri !== "string") {
  throw new Error("Connection URI is not provided");
}
