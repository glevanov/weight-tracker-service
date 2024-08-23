import { MongoClient } from "mongodb";

import { config } from "./config.js";
import type { ErrorResult, Weight, Result } from "./types.js";
import { literals } from "./literals.js";
import { getSalt, hashPassword, verifyPassword } from "./auth/auth.js";

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
      error: err,
    };
  }

  async addWeight(weight: number): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error(literals.error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.weightsCollection);

      await collection.insertOne({
        weight,
        timestamp: new Date(),
        user: "glevanov",
      });

      await this.#connection.close();

      return {
        isSuccess: true,
        data: literals.response.weight.addSuccess,
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async getWeights(start: Date, end: Date): Promise<Result<Weight[]>> {
    try {
      if (this.#connection === null) {
        throw new Error(literals.error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.weightsCollection);

      const weights = await collection
        .find({
          user: "glevanov",
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
  ): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error(literals.error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.usersCollection);

      const existingUser = await collection.findOne({ username });
      if (existingUser) {
        await this.#connection.close();
        return {
          isSuccess: false,
          error: new Error(literals.error.user.exists),
        };
      }

      const salt = getSalt();
      const hashedPassword = await hashPassword(password, salt);

      if (typeof hashedPassword !== "string") {
        throw new Error(literals.error.user.hashFailed);
      }

      await collection.insertOne({
        username,
        password: hashedPassword,
        salt,
      });

      await this.#connection.close();

      return {
        isSuccess: true,
        data: literals.response.user.registerSuccess,
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async loginUser(username: string, password: string): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error(literals.error.connection.notSet);
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.usersCollection);

      const user = await collection.findOne({ username });
      if (!user) {
        await this.#connection.close();
        return {
          isSuccess: false,
          error: new Error(literals.error.user.unauthorized),
        };
      }

      const hashedPassword = await hashPassword(password, user.salt);
      const isPasswordCorrect =
        typeof hashedPassword === "string" &&
        (await verifyPassword(hashedPassword, user.salt, user.password)) ===
          true;
      if (!isPasswordCorrect) {
        await this.#connection.close();
        return {
          isSuccess: false,
          error: new Error(literals.error.user.unauthorized),
        };
      }

      // todo sessions

      await this.#connection.close();

      return {
        isSuccess: true,
        data: "User is authorized",
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }
}
