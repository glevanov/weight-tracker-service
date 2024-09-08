import { randomBytes, scrypt } from "node:crypto";

export const hashPassword = (
  password: string,
  salt: string,
): Promise<string | Error> => {
  return new Promise((resolve, reject) => {
    scrypt(password, salt, 64, (err, derivedKey) => {
      if (err) reject(err);
      resolve(derivedKey.toString("hex"));
    });
  });
};

export const generateSalt = () => randomBytes(16).toString("hex");
