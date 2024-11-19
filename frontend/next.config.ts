import { NextConfig } from "next";

const nextConfig: NextConfig = {
    output: "standalone",
    async rewrites() {
        return [
            {
                source: "/api/:path*",
                destination: `http://${process.env.NEXT_PUBLIC_API_URL}/api/:path*`,
            },
        ];
    },
};

export default nextConfig;
