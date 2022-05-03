/** @type {import('next').NextConfig} */
const removeImports = require('next-remove-imports')();
module.exports = removeImports({
    // reactStrictMode: true,
    // pageExtensions: ['page.tsx', 'page.ts'],
    images: {
      domains: [],
    },
    swcMinify: true,
    i18n: {
      locales: ["en", "ja"],
      defaultLocale: "ja",
    },
});
