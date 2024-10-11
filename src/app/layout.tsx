import AuthProvider from "@/providers/AuthProvider";
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
    title: "Social Network",
    description: "",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body>
                <AuthProvider>{children}</AuthProvider>
            </body>
        </html>
    );
}
