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
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <div className="shadow-xl w-full mb-5 p-3">
                        <h2 className="text-black text-xl font-bold font-['Inter'] tracking-wide text-center">Settings</h2>
                    </div>
                    <UserInfos {...data} />
                </div>
            </HomeProfileLayout>
        </>
    );
}
