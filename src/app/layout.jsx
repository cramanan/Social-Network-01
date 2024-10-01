import UserContextProvider from "@/providers/UserContextProvider";
import "./globals.css";

export const metadata = {
    title: "Social network",
    description: "",
};

export default function RootLayout({ children }) {
    return (
        <html lang="en">
            <UserContextProvider>
                <body>{children}</body>
            </UserContextProvider>
        </html>
    );
}
