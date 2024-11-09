import { EditableUser } from "@/types/user";
import React from "react";
import UserInfos from "./UserInfos";
import { headers as requestHeaders } from "next/headers";

export default async function Page() {
    const headers = await requestHeaders();
    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/auth`,
        {
            headers,
            cache: "no-cache",
        }
    );
    const data: EditableUser = await response.json();

    return <UserInfos {...data} />;
}
