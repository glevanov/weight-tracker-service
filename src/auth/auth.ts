import { randomBytes, scrypt } from "node:crypto";

export const hashPassword = (password: string, salt: string) => {
  return new Promise((resolve, reject) => {
    scrypt(password, salt, 64, (err, derivedKey) => {
      if (err) reject(err);
      resolve(derivedKey.toString("hex"));
    });
  });
};

export const verifyPassword = (
  password: string,
  salt: string,
  hashedPassword: string,
) => {
  return new Promise((resolve, reject) => {
    scrypt(password, salt, 64, (err, derivedKey) => {
      if (err) reject(err);
      resolve(derivedKey.toString("hex") === hashedPassword);
    });
  });
};

export const getSalt = () => randomBytes(16).toString("hex");
