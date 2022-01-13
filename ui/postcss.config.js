// module.exports = {
//   plugins: {
//     tailwindcss: {},
//     "vue-cli-plugin-tailwind/purgecss": {},
//     autoprefixer: {}
//   }
// }

module.exports = {
  plugins: [
    // ...
    require("tailwindcss"),
    require("autoprefixer")
    // ...
  ]
};
