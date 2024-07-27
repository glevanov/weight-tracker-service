import { MongoClient } from "mongodb";

import { config } from "../config.mjs";

export class Connection {
  /**
   * The instance of the Connection
   * @type {Connection}
   * @private
   */
  static #instance = null;

  /**
   * The connection to the database
   * @type {MongoClient}
   * @private
   */
  #connection = null;

  constructor() {}

  /**
   * Get the instance of the connection
   * @returns {Connection}
   */
  static get instance() {
    if (!Connection.#instance) {
      Connection.#instance = new Connection();
      Connection.#instance.#setConnection();
    }

    return Connection.#instance;
  }

  /**
   * Set the connection to the database
   * @private
   */
  #setConnection() {
    this.#connection = new MongoClient(config.connectionUri);
  }

  /**
   * Close the connection and return error result
   * @param {Error} err
   * @returns {Promise<{ isSuccess: false, error: Error }>}
   */
  async #handleError(err) {
    await this.#connection.close();

    return {
      isSuccess: false,
      error: err,
    };
  }

  /**
   * Migrates all weights from integers to doubles, if any
   * @returns {Promise<{ isSuccess: true, data: null } | { isSuccess: false, error: Error }>}
   */
  async migrateWeightsIntsToDouble() {
    try {
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.collectionName);

      await collection.updateMany({ weight: { $type: "int" } }, [
        { $set: { weight: { $toDouble: "$weight" } } },
      ]);

      await this.#connection.close();

      return {
        isSuccess: true,
        data: null,
      };
    } catch (err) {
      return await this.#handleError(err);
    }
  }

  /**
   * Add a new weight
   * @param {number} weight
   * @returns {Promise<{ isSuccess: true, data: null } | { isSuccess: false, error: Error }>}
   */
  async addWeight(weight) {
    try {
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
        data: null,
      };
    } catch (err) {
      return await this.#handleError(err);
    }
  }

  /**
   * Get all weights
   * @returns {Promise<{ isSuccess: true, data: { weight: number, timestamp: string }[] } | { isSuccess: false, error: Error }>}
   */
  async getWeights() {
    try {
      await this.#connection.connect();

      const db = this.#connection.db(config.dbName);
      const collection = db.collection(config.collectionName);

      const weights = await collection.find().toArray();

      await this.#connection.close();

      return {
        isSuccess: true,
        data: weights.map(({ weight, timestamp }) => ({ weight, timestamp })),
      };
    } catch (err) {
      return await this.#handleError(err);
    }
  }
}
