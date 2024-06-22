/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'gray-900': '#1a1a1a',
        'gray-800': '#2a2a2a',
        'gray-700': '#3a3a3a',
        'gray-600': '#4a4a4a',
        'blue-500': '#3b82f6',
        'blue-400': '#60a5fa',
      },
    },
  },
  plugins: [],
};
