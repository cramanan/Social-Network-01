import { createContext, useContext } from "react";
import { User } from "@/types/user";

type ContextType = {
    user: User | undefined;
    setUser: (user: User) => void;
};

export const authContext = createContext<ContextType>({
    user: undefined,
    setUser: () => {},
});

export const useAuth = () => useContext(authContext);
