var en = require("./translations.en.json");
var ja = require("./translations.ja.json");
var fr = require("./translations.fr.json");

const i18n = {
  translations: {
    en,
    ja,
    fr,
  },
  defaultLang: "en",
  useBrowserDefault: true,
};

module.exports = i18n;
