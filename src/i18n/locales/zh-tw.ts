import type { Locale } from "../i18n.js";

export const zhTW: Locale = {
  validation: {
    weight: {
      failedToParse: "無法解析重量格式",
      invalidWeightNumberFormat:
        "無法識別重量格式。預期是一個兩位或三位數，選擇性地帶有小數部分。",
      weightTooLow: "重量不能少於 10 公斤。",
      weightTooHigh: "重量不能超過 200 公斤。",
    },
    timestamp: {
      notString: "時間戳記不是字串。",
      notDate: "無法識別日期格式。",
      failedToParseStart: "無法解析開始日期格式",
      failedToParseEnd: "無法解析結束日期格式",
    },
    auth: {
      failedToParse: "無法解析使用者數據格式",
      invalidFormat: "使用者名稱或密碼不是字串",
    },
  },
  response: {
    weight: {
      addSuccess: "重量新增成功",
    },
    migration: {
      success: "遷移成功完成",
    },
    user: {
      registerSuccess: "使用者註冊成功",
    },
  },
  error: {
    connection: {
      notSet: "尚未建立資料庫連線",
    },
    user: {
      exists: "此名稱的使用者已經存在",
      hashFailed: "無法加密密碼",
      failedToAuthorize: "授權失敗",
      unauthorized: "使用者未經授權",
    },
    unknown: "發生未知錯誤",
  },
};
