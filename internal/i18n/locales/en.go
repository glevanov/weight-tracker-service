package locales

var En = Locale{
	Validation: ValidationLocale{
		Weight: WeightValidationLocale{
			FailedToParse:             "Failed to parse weight format",
			InvalidWeightNumberFormat: "Failed to recognize weight format. Expected a two- or three-digit number, optionally with a decimal part.",
			WeightTooLow:              "Weight cannot be less than 10 kg.",
			WeightTooHigh:             "Weight cannot be more than 200 kg.",
		},
		Timestamp: TimestampValidationLocale{
			NotString:          "Timestamp is not a string.",
			NotDate:            "Failed to recognize date format.",
			FailedToParseStart: "Failed to parse start date format",
			FailedToParseEnd:   "Failed to parse end date format",
		},
		Auth: AuthValidationLocale{
			FailedToParse: "Failed to parse user data format",
			InvalidFormat: "Username or password is not a string",
		},
	},
	Response: ResponseLocale{
		Weight:    "Weight added successfully",
		Migration: "Migration completed successfully",
		User:      "User registered successfully",
	},
	Error: ErrorLocale{
		Connection: ConnectionErrorLocale{
			NotSet: "Database connection not established",
		},
		User: UserErrorLocale{
			Exists:            "User with this name already exists",
			HashFailed:        "Failed to hash the password",
			FailedToAuthorize: "Authorization failed",
			Unauthorized:      "User is not authorized",
		},
		Unknown: "An unknown error occurred",
	},
}
