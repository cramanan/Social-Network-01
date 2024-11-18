"use client";

import { GroupList } from "@/components/GroupList";
import NewGroup from "./NewGroup";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default function Page() {
    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <NewGroup />
                    <GroupList />
                </div>
            </HomeProfileLayout>
        </>
    );
}
