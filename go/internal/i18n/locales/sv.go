package locales

var Sv = Locale{
	Validation: ValidationLocale{
		Weight: WeightValidationLocale{
			FailedToParse:             "Misslyckades med att tolka viktformatet",
			InvalidWeightNumberFormat: "Kunde inte känna igen viktformatet. Förväntades ett två- eller tresiffrigt tal, eventuellt med en decimaldel.",
			WeightTooLow:              "Vikten kan inte vara mindre än 10 kg.",
			WeightTooHigh:             "Vikten kan inte vara mer än 200 kg.",
		},
		Timestamp: TimestampValidationLocale{
			NotString:          "Tidsstämpeln är inte en sträng.",
			NotDate:            "Kunde inte känna igen datumformatet.",
			FailedToParseStart: "Misslyckades med att tolka startdatumformatet",
			FailedToParseEnd:   "Misslyckades med att tolka slutdatumformatet",
		},
		Auth: AuthValidationLocale{
			FailedToParse: "Misslyckades med att tolka användardataformatet",
			InvalidFormat: "Användarnamn eller lösenord är inte en sträng",
		},
	},
	Response: ResponseLocale{
		Weight:    "Vikt tillagd framgångsrikt",
		Migration: "Migrering slutförd framgångsrikt",
		User:      "Användare registrerad framgångsrikt",
	},
	Error: ErrorLocale{
		Connection: ConnectionErrorLocale{
			NotSet: "Databasanslutning inte etablerad",
		},
		User: UserErrorLocale{
			Exists:            "Användare med detta namn finns redan",
			HashFailed:        "Misslyckades med att hasha lösenordet",
			FailedToAuthorize: "Auktorisering misslyckades",
			Unauthorized:      "Användaren är inte auktoriserad",
		},
		Unknown: "Ett okänt fel uppstod",
	},
}
