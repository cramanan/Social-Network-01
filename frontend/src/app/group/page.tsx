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

    const handlersClick = [handleJoinedGroupsClick, handleAllGroupsClick]

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <NewGroup />
                    <ul className="flex flex-row w-full justify-evenly">
                        {["Joined groups", "All groups"].map((name, idx) => (
                            <li key={idx} className="w-1/2 text-center cursor-pointer p-3 hover:bg-black/10" onClick={handlersClick[idx]}>{name}</li>
                        ))}
                    </ul>

                    {showJoinedGroups && (
                        <span>Joined Groups</span>
                    )}

                    {showAllGroups && (
                        <GroupList />
                    )}
                </div>
            </HomeProfileLayout>
        </>
    );
}
