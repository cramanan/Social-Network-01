"use client";

import { GroupList } from "@/components/GroupList";
import NewGroup from "./NewGroup";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useState } from "react";

export default function Page() {
    const [showAllGroups, setShowAllGroups] = useState(false)
    const [showJoinedGroups, setShowJoinedGroups] = useState(true)

    const handleAllGroupsClick = () => {
        if (!showAllGroups) {
            setShowAllGroups(!showAllGroups)
            setShowJoinedGroups(false)
        }
    }

    const handleJoinedGroupsClick = () => {
        if (!showJoinedGroups) {
            setShowAllGroups(false)
            setShowJoinedGroups(!showJoinedGroups)
        }
    }


    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <NewGroup />
                    <ul className="flex flex-row w-full justify-evenly p-5">
                        <li onClick={handleJoinedGroupsClick}>Joined groups</li>
                        <li onClick={handleAllGroupsClick}>All groups</li>
                    </ul>
                    {showAllGroups && (
                        <GroupList />
                    )}

                    {showJoinedGroups && (
                        <span>Joined Groups</span>
                    )}

                </div>
            </HomeProfileLayout>
        </>
    );
}
