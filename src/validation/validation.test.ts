import { expect, test } from "vitest";
import {
  validateAndFormatWeight,
  WeightValidationError,
} from "./validation.js";

test("validateAndFormatWeight", () => {
  expect(validateAndFormatWeight("82", "en")).toBe(82);
  expect(validateAndFormatWeight("82,2", "en")).toBe(82.2);
  expect(validateAndFormatWeight("82.2", "en")).toBe(82.2);
  expect(validateAndFormatWeight("83,40", "en")).toBe(83.4);
  expect(validateAndFormatWeight("83.40", "en")).toBe(83.4);
  expect(validateAndFormatWeight("83,405", "en")).toBe(83.41);
  expect(validateAndFormatWeight("83.405", "en")).toBe(83.41);
  expect(validateAndFormatWeight("83,404", "en")).toBe(83.4);
  expect(validateAndFormatWeight("83.404", "en")).toBe(83.4);
  expect(validateAndFormatWeight("103", "en")).toBe(103);
  expect(validateAndFormatWeight("103.4", "en")).toBe(103.4);
  expect(validateAndFormatWeight("103.405", "en")).toBe(103.41);
  expect(validateAndFormatWeight("199.993", "en")).toBe(199.99);

  expect(validateAndFormatWeight("5", "en")).toBeInstanceOf(
    WeightValidationError,
  );
  expect(validateAndFormatWeight("9", "en")).toBeInstanceOf(
    WeightValidationError,
  );
  expect(validateAndFormatWeight("500", "en")).toBeInstanceOf(
    WeightValidationError,
  );
  expect(validateAndFormatWeight("250", "en")).toBeInstanceOf(
    WeightValidationError,
  );
  expect(validateAndFormatWeight("44.1.1", "en")).toBeInstanceOf(
    WeightValidationError,
  );
});
