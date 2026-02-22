package i18n

import (
	"net/http"
	"strings"

	"weight-tracker-service/internal/i18n/locales"
)

var supportedLocales = map[string]locales.Locale{
	"en": locales.En,
	"ru": locales.Ru,
	"sv": locales.Sv,
}

func ExtractLang(r *http.Request) string {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		return "en"
	}

	lang = strings.Split(lang, ",")[0]
	lang = strings.Split(lang, "-")[0]

	if _, ok := supportedLocales[lang]; ok {
		return lang
	}

	return "en"
}

func GetLocale(lang string) locales.Locale {
	if locale, ok := supportedLocales[lang]; ok {
		return locale
	}
	return locales.En
}

func Translate(lang, key string) string {
	locale := GetLocale(lang)

	switch key {
	case "validation.weight.failedToParse":
		return locale.Validation.Weight.FailedToParse
	case "validation.weight.invalidWeightNumberFormat":
		return locale.Validation.Weight.InvalidWeightNumberFormat
	case "validation.weight.weightTooLow":
		return locale.Validation.Weight.WeightTooLow
	case "validation.weight.weightTooHigh":
		return locale.Validation.Weight.WeightTooHigh
	case "validation.timestamp.notString":
		return locale.Validation.Timestamp.NotString
	case "validation.timestamp.notDate":
		return locale.Validation.Timestamp.NotDate
	case "validation.timestamp.failedToParseStart":
		return locale.Validation.Timestamp.FailedToParseStart
	case "validation.timestamp.failedToParseEnd":
		return locale.Validation.Timestamp.FailedToParseEnd
	case "validation.auth.failedToParse":
		return locale.Validation.Auth.FailedToParse
	case "validation.auth.invalidFormat":
		return locale.Validation.Auth.InvalidFormat
	case "response.weight.addSuccess":
		return locale.Response.Weight
	case "response.migration.success":
		return locale.Response.Migration
	case "response.user.registerSuccess":
		return locale.Response.User
	case "error.connection.notSet":
		return locale.Error.Connection.NotSet
	case "error.user.exists":
		return locale.Error.User.Exists
	case "error.user.hashFailed":
		return locale.Error.User.HashFailed
	case "error.user.failedToAuthorize":
		return locale.Error.User.FailedToAuthorize
	case "error.user.unauthorized":
		return locale.Error.User.Unauthorized
	case "error.unknown":
		return locale.Error.Unknown
	default:
		return locale.Error.Unknown
	}
}
