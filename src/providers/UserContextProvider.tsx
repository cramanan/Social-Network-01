import { ReactNode } from "react";
import { userContext } from "./UserContext";

export default function UserContextProvider({
    children,
}: {
    children: ReactNode;
}) {
    return <userContext.Provider value={null}>{children}</userContext.Provider>;
}
