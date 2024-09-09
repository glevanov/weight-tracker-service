import { Lang, locales } from "../i18n/i18n.js";

export class WeightValidationError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "WeightValidationError";
  }
}

export const validateAndFormatWeight = (input: string, lang: Lang) => {
  const withUniformSeparator = input.trim().replace(",", ".");

  const regex = /^\d+(\.\d+)?$/;
  if (!regex.test(withUniformSeparator)) {
    return new WeightValidationError(
      locales[lang].validation.weight.invalidWeightNumberFormat,
    );
  }

  const parsed = Number(Number.parseFloat(withUniformSeparator).toFixed(2));

  if (parsed < 10) {
    return new WeightValidationError(
      locales[lang].validation.weight.weightTooLow,
    );
  }
  if (parsed > 200) {
    return new WeightValidationError(
      locales[lang].validation.weight.weightTooHigh,
    );
  }

  return parsed;
};

export class TimestampValidationError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "TimestampValidationError";
  }
}

export const validateAndParseTimestamp = (input: unknown, lang: Lang) => {
  if (typeof input !== "string") {
    return new TimestampValidationError(
      locales[lang].validation.timestamp.notString,
    );
  }

  const date = new Date(input);

  if (isNaN(date.getTime())) {
    return new TimestampValidationError(
      locales[lang].validation.timestamp.notDate,
    );
  }

  return date;
};
