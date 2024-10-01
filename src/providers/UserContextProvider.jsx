"use client";

import UserContext from "./UserContext";

export default function UserContextProvider({ children }) {
    return <UserContext.Provider>{children}</UserContext.Provider>;
}
