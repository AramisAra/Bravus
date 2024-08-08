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
        'blue-600': '#2563eb',
        'indigo-600': '#4f46e5',
        'Lavender': '#E6E6FA',
        'Medium Purple':'#9370DB',
        'Purple': '#800080',
        'Dark Purple': '#4B0082',
        'Black': '#000000',
        'Charcoal': '#36454F',
        'Jet Black': '#343434',
      },
    },
  },
  plugins: [],
};
