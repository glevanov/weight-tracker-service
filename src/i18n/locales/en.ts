import type { Locale } from "../i18n.js";

export const en: Locale = {
  validation: {
    weight: {
      failedToParse: "Failed to parse weight format",
      invalidWeightNumberFormat:
        "Failed to recognize weight format. Expected a two- or three-digit number, optionally with a decimal part.",
      weightTooLow: "Weight cannot be less than 10 kg.",
      weightTooHigh: "Weight cannot be more than 200 kg.",
    },
    timestamp: {
      notString: "Timestamp is not a string.",
      notDate: "Failed to recognize date format.",
      failedToParseStart: "Failed to parse start date format",
      failedToParseEnd: "Failed to parse end date format",
    },
    auth: {
      failedToParse: "Failed to parse user data format",
      invalidFormat: "Username or password is not a string",
    },
  },
  response: {
    weight: {
      addSuccess: "Weight added successfully",
    },
    migration: {
      success: "Migration completed successfully",
    },
    user: {
      registerSuccess: "User registered successfully",
    },
  },
  error: {
    connection: {
      notSet: "Database connection not established",
    },
    user: {
      exists: "User with this name already exists",
      hashFailed: "Failed to hash the password",
      failedToAuthorize: "Authorization failed",
      unauthorized: "User is not authorized",
    },
    unknown: "An unknown error occurred",
  },
};
