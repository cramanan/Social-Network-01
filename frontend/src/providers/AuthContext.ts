import { createContext } from "react";
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
    ) => Promise<void>;
    login: (email: string, password: string) => Promise<void>;
    logout: () => Promise<void>;
};

export const authContext = createContext<ContextType>({
    user: null,
    loading: false,
    signup: async () => {},
    login: async () => {},
    logout: async () => {},
});
