import { literals } from "./literals.js";

export class WeightValidationError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "WeightValidationError";
  }
}

export const validateAndFormatWeight = (input: string) => {
  const withUniformSeparator = input.trim().replace(",", ".");

  const regex = /^\d+(\.\d+)?$/;
  if (!regex.test(withUniformSeparator)) {
    return new WeightValidationError(
      literals.validation.weight.invalidWeightNumberFormat,
    );
  }

  const parsed = Number(Number.parseFloat(withUniformSeparator).toFixed(2));

  if (parsed < 10) {
    return new WeightValidationError(literals.validation.weight.weightTooLow);
  }
  if (parsed > 200) {
    return new WeightValidationError(literals.validation.weight.weightTooHigh);
  }

  return parsed;
};

export class TimestampValidationError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "TimestampValidationError";
  }
}

export const validateAndParseTimestamp = (input: unknown) => {
  if (typeof input !== "string") {
    return new TimestampValidationError(
      literals.validation.timestamp.notString,
    );
  }

  const date = new Date(input);

  if (isNaN(date.getTime())) {
    return new TimestampValidationError(literals.validation.timestamp.notDate);
  }

  return date;
};
