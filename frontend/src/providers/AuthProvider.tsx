"use client";

import { ReactNode, useEffect, useState } from "react";
import { authContext } from "./AuthContext";
import { User } from "@/types/user";

export default function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    const login = async (email: string, password: string) => {
        const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_URL}/api/login`,
            {
                headers: { "Content-Type": "application/json" },
                method: "POST",
                body: JSON.stringify({ email, password }),
            }
        );
        const data = await response.json();
        setUser(data);
    };

    const signup = async (
        nickname: string,
        email: string,
        password: string,
        firstName: string,
        lastName: string,
        dateOfBirth: string
    ) => {
        const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_URL}/api/register`,
            {
                headers: { "Content-Type": "application/json" },
                method: "POST",
                body: JSON.stringify({
                    nickname,
                    email,
                    password,
                    firstName,
                    lastName,
                    dateOfBirth,
                }),
            }
        );
        const data = await response.json();
        setUser(data);
    };

    const logout = async () => {
        await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/logout`);
        setUser(null);
        setLoading(false);
    };

    useEffect(() => {
        const authenticate = async () => {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/api/auth`
            );
            const data = await response.json();
            setUser(data);
            setLoading(false);
        };
        authenticate();
    }, []);

    return (
        <authContext.Provider value={{ user, loading, signup, login, logout }}>
            {children}
        </authContext.Provider>
    );
}
