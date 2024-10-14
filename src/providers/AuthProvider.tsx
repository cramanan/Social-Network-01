"use client";

import { ReactNode, useEffect, useState } from "react";
import { authContext } from "./AuthContext";
import { User } from "@/types/user";

export default function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        fetch("/api/auth")
            .then((resp) => (resp.ok ? resp.json() : null))
            .then(setUser)
            .then(() => console.log("calling"))
            .catch(console.error);
    }, []);

    return (
        <authContext.Provider value={{ user, setUser }}>
            {children}
        </authContext.Provider>
    );
}
