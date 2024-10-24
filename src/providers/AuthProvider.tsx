"use client";

import { ReactNode, useEffect, useState } from "react";
import { authContext } from "./AuthContext";
import { User } from "@/types/user";

export default function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | undefined>(undefined);

    useEffect(() => {
        fetch("/api/auth")
            .then((resp) => (resp.ok ? resp.json() : null))
            .then(setUser)
            .catch(console.error);
    }, []);

    return (
        <authContext.Provider value={{ user, setUser }}>
            {children}
        </authContext.Provider>
    );
}
