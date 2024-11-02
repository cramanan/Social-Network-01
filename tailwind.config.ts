import type { Config } from "tailwindcss";


const config: Config = {
    content: [
        "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/layouts/**/*.{js,ts,jsx,tsx,mdx}",
    ],

    plugins: [],
    safelist: [
    'bg-red-500',
    'bg-green-500'
  ]
};
export default config;
