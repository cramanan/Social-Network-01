"use client";

import { useUser } from "@/providers/UserContext";
import Link from "next/link";
import { redirect } from "next/navigation";
import { useContext } from "react";

export default function Layout({ children }: { children: React.ReactNode }) {
    const user = useUser();
    if (!user) return redirect("/auth");

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
