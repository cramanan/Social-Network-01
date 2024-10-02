/** @type {import('next').NextConfig} */
const nextConfig = {
    async rewrites() {
        return [
            {
                source: "/api/:path*",
                destination: `${process.env.API_URL}/api/:path*`, // Proxy to Backend
            },
        ];
    },
};

export default nextConfig;
