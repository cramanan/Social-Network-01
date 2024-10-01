"use client";

import UserContext from "@/providers/UserContext";
import Image from "next/image";
import Link from "next/link";
import { useContext } from "react";

export default function Header() {
    return (
        <header>
            <Link href="/">SocialNetwork</Link>
            <input type="text" />
            <div>
                <Image
                    src=""
                    width={40}
                    height={40}
                    alt="Your profile picture"
                />
                <div>John</div>
                <div>Doe</div>
            </div>
        </header>
    );
}
