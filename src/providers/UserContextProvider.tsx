import { ReactNode } from "react";
import UserContext from "./UserContext";

export default function UserContextProvider({
    children,
}: {
    children: ReactNode;
}) {
    return <UserContext.Provider value={null}>{children}</UserContext.Provider>;
}
