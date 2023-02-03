/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  eslint: {
    dirs: ["pages", "types", "components"], // Only run ESLint on the 'pages' and 'utils' directories during production builds (next build)
  },
  env: {
    GOOGLE_MAP_API_KEY: process.env.GOOGLE_MAP_API_KEY,
  },
};

module.exports = nextConfig;
