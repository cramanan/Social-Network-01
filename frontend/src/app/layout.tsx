import AuthProvider from "@/providers/AuthProvider";
import type { Metadata } from "next";
import "./globals.css";
import WebSocketProvider from "@/providers/WebSocketProvider";

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
                <AuthProvider>
                    <WebSocketProvider>{children}</WebSocketProvider>
                </AuthProvider>
            </body>
        </html>
    );
}
