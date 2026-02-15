package locales

type Locale struct {
	Validation ValidationLocale
	Response   ResponseLocale
	Error      ErrorLocale
}

type ValidationLocale struct {
	Weight    WeightValidationLocale
	Timestamp TimestampValidationLocale
	Auth      AuthValidationLocale
}

type WeightValidationLocale struct {
	FailedToParse             string
	InvalidWeightNumberFormat string
	WeightTooLow              string
	WeightTooHigh             string
}

type TimestampValidationLocale struct {
	NotString          string
	NotDate            string
	FailedToParseStart string
	FailedToParseEnd   string
}

type AuthValidationLocale struct {
	FailedToParse string
	InvalidFormat string
}

type ResponseLocale struct {
	Weight    string
	Migration string
	User      string
}

type ErrorLocale struct {
	Connection ConnectionErrorLocale
	User       UserErrorLocale
	Unknown    string
}

type ConnectionErrorLocale struct {
	NotSet string
}

type UserErrorLocale struct {
	Exists            string
	HashFailed        string
	FailedToAuthorize string
	Unauthorized      string
}
