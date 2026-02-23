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

var translationsCache = map[string]map[string]string{}

func init() {
	for lang, locale := range supportedLocales {
		translationsCache[lang] = map[string]string{
			"validation.weight.failedToParse":             locale.Validation.Weight.FailedToParse,
			"validation.weight.invalidWeightNumberFormat": locale.Validation.Weight.InvalidWeightNumberFormat,
			"validation.weight.weightTooLow":              locale.Validation.Weight.WeightTooLow,
			"validation.weight.weightTooHigh":             locale.Validation.Weight.WeightTooHigh,
			"validation.timestamp.notString":              locale.Validation.Timestamp.NotString,
			"validation.timestamp.notDate":                locale.Validation.Timestamp.NotDate,
			"validation.timestamp.failedToParseStart":     locale.Validation.Timestamp.FailedToParseStart,
			"validation.timestamp.failedToParseEnd":       locale.Validation.Timestamp.FailedToParseEnd,
			"validation.auth.failedToParse":               locale.Validation.Auth.FailedToParse,
			"validation.auth.invalidFormat":               locale.Validation.Auth.InvalidFormat,
			"response.weight.addSuccess":                  locale.Response.Weight,
			"response.migration.success":                  locale.Response.Migration,
			"response.user.registerSuccess":               locale.Response.User,
			"error.connection.notSet":                     locale.Error.Connection.NotSet,
			"error.user.exists":                           locale.Error.User.Exists,
			"error.user.hashFailed":                       locale.Error.User.HashFailed,
			"error.user.failedToAuthorize":                locale.Error.User.FailedToAuthorize,
			"error.user.unauthorized":                     locale.Error.User.Unauthorized,
			"error.unknown":                               locale.Error.Unknown,
		}
	}
}

func Translate(lang, key string) string {
	if t, ok := translationsCache[lang][key]; ok {
		return t
	}
	return translationsCache["en"]["error.unknown"]
}
