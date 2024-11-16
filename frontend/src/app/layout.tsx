import AuthProvider from "@/providers/AuthProvider";
import type { Metadata } from "next";
import "./globals.css";
import WebSocketProvider from "@/providers/WebSocketProvider";
import NotificationLayout from "@/layouts/NotificationLayout";

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
                    <WebSocketProvider>
                        <NotificationLayout>{children}</NotificationLayout>
                    </WebSocketProvider>
                </AuthProvider>
            </body>
        </html>
    );
}
