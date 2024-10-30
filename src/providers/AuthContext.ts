import { createContext, useContext } from "react";
import { User } from "@/types/user";

type ContextType = {
    user: User | undefined;
    loading: boolean;
    setUser: (user: User) => void;
};

export const authContext = createContext<ContextType>({
    user: undefined,
    loading: true,
    setUser: () => {},
});

export const useAuth = () => useContext(authContext);
