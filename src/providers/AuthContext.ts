import { createContext, useContext } from "react";
import { User } from "@/types/user";

type ContextType = {
    user: User | null;
    loading: boolean;
    signup: (
        nickname: string,
        email: string,
        password: string,
        firstName: string,
        lastName: string,
        dateOfBirth: string
    ) => void;
    login: (email: string, password: string) => void;
    logout: () => void;
};

export const authContext = createContext<ContextType>({
    user: null,
    loading: false,
    signup: () => {},
    login: () => {},
    logout: () => {},
});

export const useAuth = () => useContext(authContext);
