"use client";

import UserContext from "@/providers/UserContext";
import Link from "next/link";
import { useContext } from "react";

export default function Layout({ children }: { children: React.ReactNode }) {
    const user = useContext(UserContext);
    return (
        <>
            <header>
                <Link href="/">SocialNetwork</Link>
                <div>{JSON.stringify(user)}</div>
            </header>
            {children}
            <footer></footer>
        </>
    );
}
