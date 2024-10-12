"use client";

import Link from "next/link";

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
