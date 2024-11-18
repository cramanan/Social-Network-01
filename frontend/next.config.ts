import { NextConfig } from "next";

const nextConfig: NextConfig = {
    reactStrictMode: false,
    async rewrites() {
        return [
            {
                source: "/api/:path*",
                destination: `${process.env.NEXT_PUBLIC_API_URL}/api/:path*`,
            },
        ];
    },
};

export default nextConfig;
