/** @type{import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./static/**/*.css", "./*.html"],
  safelist: [
    "font-mono2", // Force Tailwind to generate this class
  ],
  theme: {
    extend: {
      fontFamily: {
        mono2: ["JetBrains Mono", "monospace"],
      },
    },
  },
  plugins: [],
};
