import { createContext } from "react";
import { User } from "@/types/user";

export default createContext<null | User>(null);
