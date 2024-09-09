import type { Locale } from "../i18n.js";

export const sv: Locale = {
  validation: {
    weight: {
      failedToParse: "Misslyckades med att tolka viktformatet",
      invalidWeightNumberFormat:
        "Kunde inte känna igen viktformatet. Förväntade ett två- eller tresiffrigt nummer, eventuellt med decimaldel.",
      weightTooLow: "Vikten kan inte vara mindre än 10 kg.",
      weightTooHigh: "Vikten kan inte vara mer än 200 kg.",
    },
    timestamp: {
      notString: "Tidsstämpeln är inte en sträng.",
      notDate: "Kunde inte känna igen datumformatet.",
      failedToParseStart: "Misslyckades med att tolka startdatumformatet",
      failedToParseEnd: "Misslyckades med att tolka slutdatumformatet",
    },
    auth: {
      failedToParse: "Misslyckades med att tolka användardataformatet",
      invalidFormat: "Användarnamn eller lösenord är inte en sträng",
    },
  },
  response: {
    weight: {
      addSuccess: "Vikt tillagd framgångsrikt",
    },
    migration: {
      success: "Migrering slutförd framgångsrikt",
    },
    user: {
      registerSuccess: "Användare registrerad framgångsrikt",
    },
  },
  error: {
    connection: {
      notSet: "Databasanslutning inte etablerad",
    },
    user: {
      exists: "Användare med detta namn finns redan",
      hashFailed: "Misslyckades med att hasha lösenordet",
      failedToAuthorize: "Auktorisering misslyckades",
      unauthorized: "Användaren är inte auktoriserad",
    },
    unknown: "Ett okänt fel uppstod",
  },
};
