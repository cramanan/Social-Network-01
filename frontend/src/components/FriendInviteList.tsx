import React from "react";
import FriendInvite from "./FriendInvite";

const FriendInviteList = () => {
    return (
        <div
            id="friendInviteList"
            className="relative flex flex-col w-full h-full xl:w-fit xl:h-fit xl:bg-white/25 xl:rounded-[30px] xl:px-2 xl:py-5"
        >
            <h2 className="text-4xl text-white text-center py-5 xl:sr-only">Friend Request List</h2>

            <div className="flex flex-col h-[75vh] items-center gap-3 mx-2 overflow-scroll no-scrollbar xl:max-h-[68vh] xl:gap-1">
                <FriendInvite />
                <FriendInvite />
                <FriendInvite />
                <FriendInvite />
                <FriendInvite />
            </div>
        </div>
    );
};

export default FriendInviteList;
