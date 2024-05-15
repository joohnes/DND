/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/routes/**/*.{svelte,js,ts}', './src/components/*.svelte'],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
}
