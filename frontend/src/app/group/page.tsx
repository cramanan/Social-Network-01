"use client";

import { GroupList } from "@/components/GroupList";
import NewGroup from "./NewGroup";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useState } from "react";

export default function Page() {
    const [showAllGroups, setShowAllGroups] = useState(false)
    const [showJoinedGroups, setShowJoinedGroups] = useState(true)
    const [showInviteGroups, setShowInviteGroups] = useState(false)

    const handleAllGroupsClick = () => {
        if (!showAllGroups) {
            setShowAllGroups(!showAllGroups)
            setShowJoinedGroups(false)
            setShowInviteGroups(false)
        }
    }

    const handleJoinedGroupsClick = () => {
        if (!showJoinedGroups) {
            setShowAllGroups(false)
            setShowJoinedGroups(!showJoinedGroups)
            setShowInviteGroups(false)
        }
    }

    const handleInviteGroupsClick = () => {
        if (!showInviteGroups) {
            setShowAllGroups(false)
            setShowJoinedGroups(false)
            setShowInviteGroups(!showInviteGroups)
        }
    }

    const handlersClick = [handleJoinedGroupsClick, handleAllGroupsClick, handleInviteGroupsClick]

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <NewGroup />
                    <ul className="flex flex-row w-full justify-evenly">
                        {["Joined groups", "All groups", "Invite group"].map((name, idx) => (
                            <li key={idx} className="w-1/3 text-center cursor-pointer p-3 hover:bg-black/10" onClick={handlersClick[idx]}>{name}</li>
                        ))}
                    </ul>

                    {showJoinedGroups && (
                        <span>Joined Groups</span>
                    )}

                    {showAllGroups && (
                        <GroupList />
                    )}

                    {showInviteGroups && (
                        <span>Invite group</span>
                    )}
                </div>
            </HomeProfileLayout>
        </>
    );
}
