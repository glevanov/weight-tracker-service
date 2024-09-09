import { MongoClient } from "mongodb";
import jwt from "jsonwebtoken";

import { config } from "./config.js";
import type { ErrorResult, Weight, Result } from "./types.js";
import { generateSalt, hashPassword } from "./auth/auth.js";
import { Lang, locales } from "./i18n/i18n.js";

export class Connection {
  static #instance: Connection | null = null;

  #connection: MongoClient | null = null;

  constructor() {}

  static get instance() {
    if (!Connection.#instance) {
      Connection.#instance = new Connection();
      Connection.#instance.#setConnection();
    }

    return Connection.#instance;
  }

  #setConnection() {
    this.#connection = new MongoClient(config.connectionUri);
  }

  async #handleError(err: Error): Promise<ErrorResult> {
    await this.#connection?.close();

    return {
      isSuccess: false,
      error: err.message,
    };
  }

  async addWeight(
    weight: number,
    username: string,
    lang: Lang,
  ): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error(locales[lang].error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.weightsCollection);

      await collection.insertOne({
        weight,
        timestamp: new Date(),
        user: username,
      });

      await this.#connection.close();

      return {
        isSuccess: true,
        data: locales[lang].response.weight.addSuccess,
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async getWeights(
    start: Date,
    end: Date,
    username: string,
    lang: Lang,
  ): Promise<Result<Weight[]>> {
    try {
      if (this.#connection === null) {
        throw new Error(locales[lang].error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.weightsCollection);

      const weights = await collection
        .find({
          user: username,
          timestamp: {
            $gte: start,
            $lte: end,
          },
        })
        .sort({ timestamp: 1 })
        .toArray();

      await this.#connection.close();

      return {
        isSuccess: true,
        data: weights.map(({ weight, timestamp }) => ({ weight, timestamp })),
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async registerUser(
    username: string,
    password: string,
    lang: Lang,
  ): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error(locales[lang].error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.usersCollection);

      const existingUser = await collection.findOne({ username });
      if (existingUser) {
        await this.#connection.close();
        return {
          isSuccess: false,
          error: locales[lang].error.user.exists,
        };
      }

      const salt = generateSalt();
      const hashedPassword = await hashPassword(password, salt);

      if (typeof hashedPassword !== "string") {
        throw new Error(locales[lang].error.user.hashFailed);
      }

      await collection.insertOne({
        username,
        password: hashedPassword,
        salt,
      });

      await this.#connection.close();

      return {
        isSuccess: true,
        data: locales[lang].response.user.registerSuccess,
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async loginUser(
    username: string,
    password: string,
    lang: Lang,
  ): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error(locales[lang].error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.usersCollection);

      const user = await collection.findOne({ username });
      if (!user) {
        await this.#connection.close();
        return {
          isSuccess: false,
          error: locales[lang].error.user.unauthorized,
        };
      }

      const hashedPassword = await hashPassword(password, user.salt);
      if (typeof hashedPassword !== "string") {
        throw new Error(locales[lang].error.user.unauthorized);
      }
      const isPasswordCorrect = hashedPassword === user.password;
      if (!isPasswordCorrect) {
        await this.#connection.close();
        return {
          isSuccess: false,
          error: locales[lang].error.user.unauthorized,
        };
      }

      const token = jwt.sign(
        {
          username,
        },
        config.jwtSecret,
        {
          expiresIn: config.sessionDuration,
        },
      );

      await this.#connection.close();

      return {
        isSuccess: true,
        data: token,
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }
}
