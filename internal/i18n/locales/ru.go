package locales

var Ru = Locale{
	Validation: ValidationLocale{
		Weight: WeightValidationLocale{
			FailedToParse:             "Не удалось распознать формат веса",
			InvalidWeightNumberFormat: "Не удалось распознать формат веса. Ожидалось двух- или трехзначное число, необязательно с десятичной частью.",
			WeightTooLow:              "Вес не может быть меньше 10 кг.",
			WeightTooHigh:             "Вес не может быть больше 200 кг.",
		},
		Timestamp: TimestampValidationLocale{
			NotString:          "Дата не является строкой.",
			NotDate:            "Не удалось распознать формат даты.",
			FailedToParseStart: "Не удалось распознать формат даты начала",
			FailedToParseEnd:   "Не удалось распознать формат даты конца",
		},
		Auth: AuthValidationLocale{
			FailedToParse: "Не удалось распознать формат данных пользователя",
			InvalidFormat: "Имя пользователя или пароль не являются строкой",
		},
	},
	Response: ResponseLocale{
		Weight:    "Вес успешно добавлен",
		Migration: "Миграция успешно завершена",
		User:      "Пользователь успешно зарегистрирован",
	},
	Error: ErrorLocale{
		Connection: ConnectionErrorLocale{
			NotSet: "Соединение с базой не установлено",
		},
		User: UserErrorLocale{
			Exists:            "Пользователь с таким именем уже существует",
			HashFailed:        "Не удалось захешировать пароль",
			FailedToAuthorize: "Авторизация не пройдена",
			Unauthorized:      "Пользователь не авторизован",
		},
		Unknown: "Произошла неизвестная ошибка",
	},
}
