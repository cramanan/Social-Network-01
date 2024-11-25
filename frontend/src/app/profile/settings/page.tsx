import { EditableUser } from "@/types/user";
import React from "react";
import UserInfos from "./UserInfos";
import { headers as requestHeaders } from "next/headers";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default async function Page() {
    const headers = await requestHeaders();
    const response = await fetch(
        `http://${process.env.NEXT_PUBLIC_API_URL}/api/profile`,
        { headers, cache: "no-cache" }
    );

    const data: EditableUser = await response.json();

    return (
        <>
            <HomeProfileLayout>
                <UserInfos {...data} />
            </HomeProfileLayout>
        </>
    );
}
