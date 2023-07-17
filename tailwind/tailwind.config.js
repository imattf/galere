/** @type {import('tailwindcss').Config} */
module.exports = {
  // content: [ "../templates/**/*.{gohtml,html}",], // local tailwind config
  content: [ "/templates/**/*.{gohtml,html}",], // docker tailwind config
  theme: {
    extend: {},
  },
  plugins: [],
}

