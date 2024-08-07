export const literals = {
  validation: {
    weight: {
      failedToParse: "Не удалось распознать формат веса",
      invalidWeightNumberFormat:
        "Не удалось распознать формат веса. Ожидается двух- или трёхзначное число, опционально с десятичной частью.",
      weightTooLow: "Вес не может быть меньше 10 кг.",
      weightTooHigh: "Вес не может быть больше 200 кг.",
    },
    timestamp: {
      notString: "Дата не является строкой.",
      notDate: "Не удалось распознать формат даты.",
      failedToParseStart: "Не удалось распознать формат даты начала",
      failedToParseEnd: "Не удалось распознать формат даты конца",
    },
  },
};
