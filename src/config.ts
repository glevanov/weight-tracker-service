import { loadEnvFile } from "node:process";

loadEnvFile();

if (typeof process.env.CONNECTION_URI !== "string") {
  throw new Error("Connection URI is not provided");
}
if (typeof process.env.JWT_SECRET !== "string") {
  throw new Error("JWT secret is not provided");
}

const HOUR = 1000 * 60 * 60;

export const config = {
  port: Number(process.env.PORT) || 3000,
  frontendUrl: process.env.FRONTEND_URL || "http://localhost:5173",
  jwtSecret: process.env.JWT_SECRET,
  connectionUri: process.env.CONNECTION_URI,
  dbName: "weight-tracker",
  weightsCollection: "weight",
  usersCollection: "users",
  sessionDuration: HOUR * 24 * 7,
};
