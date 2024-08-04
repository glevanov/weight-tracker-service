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
      "Не удалось распознать формат веса. Ожидается двух- или трёхзначное число, опционально с десятичной частью.",
    );
  }

  const parsed = Number(Number.parseFloat(withUniformSeparator).toFixed(2));

  if (parsed < 10) {
    return new WeightValidationError("Вес не может быть меньше 10 кг.");
  }
  if (parsed > 200) {
    return new WeightValidationError("Вес не может быть больше 200 кг.");
  }

  return parsed;
};
