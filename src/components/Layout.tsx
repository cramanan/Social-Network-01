"use client";

import { useAuth } from "@/providers/AuthContext";
import Link from "next/link";
import { redirect } from "next/navigation";

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <>
            <header>
                <Link href="/">SocialNetwork</Link>
            </header>
            {children}
            <footer></footer>
        </>
    );
}
