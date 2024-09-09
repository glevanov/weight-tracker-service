import type { IncomingMessage } from "node:http";

import { ru } from "./locales/ru.js";
import { en } from "./locales/en.js";
import { sv } from "./locales/sv.js";
import { zhTW } from "./locales/zh-tw.js";

export type Lang = "ru" | "en" | "sv" | "zh-tw";

export type Locale = typeof ru;

export const locales: Record<Lang, Locale> = {
  ru,
  en,
  sv,
  "zh-tw": zhTW,
};

const supportedLocales: Set<Lang> = new Set(["ru", "en", "sv", "zh-tw"]);

export const isSupported = (lang: unknown): lang is Lang =>
  supportedLocales.has(lang as Lang);

export const extractLangFromHeader = (req: IncomingMessage): Lang => {
  const lang = req.headers["accept-language"];
  return isSupported(lang) ? lang : "en";
};
