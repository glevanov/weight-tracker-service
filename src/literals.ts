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
    auth: {
      failedToParse: "Не удалось распознать формат данных пользователя",
      invalidFormat: "Имя пользователя или пароль не являются строкой",
    },
  },
  response: {
    weight: {
      addSuccess: "Вес успешно добавлен",
    },
    migration: {
      success: "Миграция успешно завершена",
    },
    user: {
      registerSuccess: "Пользователь успешно зарегистрирован",
    },
  },
  error: {
    connection: {
      notSet: "Соединение с базой не установлено",
    },
    user: {
      exists: "Пользователь с таким именем уже существует",
      hashFailed: "Не удалось захешировать пароль",
      failedToAuthorize: "Авторизация не пройдена",
      unauthorized: "Пользователь не авторизован",
    },
  },
};
