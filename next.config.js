/** @type {import('next').NextConfig} */
const removeImports = require('next-remove-imports')();
module.exports = removeImports({
    reactStrictMode: true,
    // when you deploy to github pages, you should specify this following
    // https://wallis.dev/blog/deploying-a-next-js-app-to-github-pages
    // assetPrefix: '/umlquiz/',
    images: {
      domains: [],
    },
    swcMinify: true,
    // i18n: {
    //   locales: ["en", "ja"],
    //   defaultLocale: "en",
    // },
});
