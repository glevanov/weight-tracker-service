import { MongoClient } from "mongodb";

import { config } from "./config.js";
import type { ErrorResult, Weight, Result } from "./types.js";

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

  async migrateWeightsIntsToDouble(): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error("Connection is not set");
      }
      await this.#connection?.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.collectionName);

      await collection.updateMany({ weight: { $type: "int" } }, [
        { $set: { weight: { $toDouble: "$weight" } } },
      ]);

      await this.#connection?.close();

      return {
        isSuccess: true,
        data: "Migration successful",
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async addWeight(weight: number): Promise<Result<string>> {
    try {
      if (this.#connection === null) {
        throw new Error("Connection is not set");
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.collectionName);

      await collection.insertOne({
        weight,
        timestamp: new Date(),
      });

      await this.#connection.close();

      return {
        isSuccess: true,
        data: "Weight added successfully",
      };
    } catch (err) {
      return await this.#handleError(err as Error);
    }
  }

  async getWeights(start: Date, end: Date): Promise<Result<Weight[]>> {
    try {
      if (this.#connection === null) {
        throw new Error("Connection is not set");
      }
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.collectionName);

      const weights = await collection
        .find({
          timestamp: {
            $gte: start,
            $lte: end,
          },
        })
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
}
