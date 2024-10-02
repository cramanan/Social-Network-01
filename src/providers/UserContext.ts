import { createContext, useContext } from "react";
import { User } from "@/types/user";

export const userContext = createContext<null | User>(null);

export const useUser = () => useContext(userContext);
