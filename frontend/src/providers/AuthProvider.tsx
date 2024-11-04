"use client";

import { ReactNode, useEffect, useState } from "react";
import { authContext } from "./AuthContext";
import { User } from "@/types/user";

export default function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    const login = async (email: string, password: string) => {
        const response = await fetch("/api/login", {
            headers: {
                "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({ email, password }),
        });
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
        const response = await fetch("/api/register", {
            headers: {
                "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({
                nickname,
                email,
                password,
                firstName,
                lastName,
                dateOfBirth,
            }),
        });
        const data = await response.json();
        setUser(data);
    };

    const logout = async () => {
        await fetch("/api/logout");
        setUser(null);
    };

    useEffect(() => {
        const login = async () => {
            const response = await fetch("/api/auth");
            const data = await response.json();
            setUser(data);
            setLoading(false);
        };
        login();
    }, []);

    return (
        <authContext.Provider value={{ user, loading, signup, login, logout }}>
            {children}
        </authContext.Provider>
    );
}
