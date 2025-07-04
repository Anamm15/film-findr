/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'inter': ['Inter', 'sans-serif'],
      },
      colors: {
        'primary': '#3f37c9',
        'secondary': '#4361ee',
        'tertiary': '#4895ef',
      }
    },
  },
  plugins: [],
}