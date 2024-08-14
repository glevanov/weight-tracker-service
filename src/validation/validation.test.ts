import { expect, test } from "vitest";
import {
  validateAndFormatWeight,
  WeightValidationError,
} from "./validation.js";

test("validateAndFormatWeight", () => {
  expect(validateAndFormatWeight("82")).toBe(82);
  expect(validateAndFormatWeight("82,2")).toBe(82.2);
  expect(validateAndFormatWeight("82.2")).toBe(82.2);
  expect(validateAndFormatWeight("83,40")).toBe(83.4);
  expect(validateAndFormatWeight("83.40")).toBe(83.4);
  expect(validateAndFormatWeight("83,405")).toBe(83.41);
  expect(validateAndFormatWeight("83.405")).toBe(83.41);
  expect(validateAndFormatWeight("83,404")).toBe(83.4);
  expect(validateAndFormatWeight("83.404")).toBe(83.4);
  expect(validateAndFormatWeight("103")).toBe(103);
  expect(validateAndFormatWeight("103.4")).toBe(103.4);
  expect(validateAndFormatWeight("103.405")).toBe(103.41);
  expect(validateAndFormatWeight("199.993")).toBe(199.99);

  expect(validateAndFormatWeight("5")).toBeInstanceOf(WeightValidationError);
  expect(validateAndFormatWeight("9")).toBeInstanceOf(WeightValidationError);
  expect(validateAndFormatWeight("500")).toBeInstanceOf(WeightValidationError);
  expect(validateAndFormatWeight("250")).toBeInstanceOf(WeightValidationError);
  expect(validateAndFormatWeight("44.1.1")).toBeInstanceOf(
    WeightValidationError,
  );
});
