package validation

const (
	ErrWeightFailedToParse         = "validation.weight.failedToParse"
	ErrWeightInvalidFormat         = "validation.weight.invalidWeightNumberFormat"
	ErrWeightTooLow                = "validation.weight.weightTooLow"
	ErrWeightTooHigh               = "validation.weight.weightTooHigh"
	ErrTimestampNotString          = "validation.timestamp.notString"
	ErrTimestampNotDate            = "validation.timestamp.notDate"
	ErrTimestampFailedToParseStart = "validation.timestamp.failedToParseStart"
	ErrTimestampFailedToParseEnd   = "validation.timestamp.failedToParseEnd"
	ErrAuthFailedToParse           = "validation.auth.failedToParse"
	ErrAuthInvalidFormat           = "validation.auth.invalidFormat"
	ErrUserUnauthorized            = "error.user.unauthorized"
	ErrUnknown                     = "error.unknown"
	ResponseWeightAdded            = "response.weight.addSuccess"
)
