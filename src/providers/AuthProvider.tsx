"use client";

import { ReactNode, useEffect, useState } from "react";
import { authContext } from "./AuthContext";
import { User } from "@/types/user";

export default function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User>();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const login = async () => {
            await fetch("/api/auth")
                .then((resp) => (resp.ok ? resp.json() : null))
                .then(setUser)
                .catch(console.error);

            setLoading(false);
        };
        login();
    }, []);

    return (
        <authContext.Provider value={{ user, loading, setUser }}>
            {children}
        </authContext.Provider>
    );
}
