import { createContext, useContext } from "react";
import { User } from "@/types/user";

interface ContextType {
    user: User | null;
    setUser: (user: User) => void;
}

export const authContext = createContext<ContextType>({
    user: null,
    setUser: () => {},
});

export const useAuth = () => useContext(authContext);
