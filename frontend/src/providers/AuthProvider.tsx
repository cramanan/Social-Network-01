"use client";

import { ReactNode, useEffect, useState } from "react";
import { authContext } from "./AuthContext";
import { User } from "@/types/user";

export default function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    const login = async (email: string, password: string) => {
        try {
            const response = await fetch("/api/login", {
                headers: { "Content-Type": "application/json" },
                method: "POST",
                body: JSON.stringify({ email, password }),
            });
            const data: User = await response.json();
            setUser(data);
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    const signup = async (
        nickname: string,
        email: string,
        password: string,
        firstName: string,
        lastName: string,
        dateOfBirth: string
    ) => {
        try {
            const response = await fetch("/api/register", {
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
            });
            const data: User = await response.json();
            setUser(data);
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    const logout = async () => {
        try {
            await fetch("/api/logout");
        } catch (error) {
            console.error(error);
        } finally {
            setUser(null);
            setLoading(false);
        }
    };

    useEffect(() => {
        const authenticate = async () => {
            try {
                const response = await fetch("/api/profile");
                const data = await response.json();
                setUser(data);
            } catch (error) {
                console.error(error);
            } finally {
                setLoading(false);
            }
        };
        authenticate();
    }, []);

    return (
        <authContext.Provider value={{ user, loading, signup, login, logout }}>
            {children}
        </authContext.Provider>
    );
}
