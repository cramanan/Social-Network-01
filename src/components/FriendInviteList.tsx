import React from "react";
import FriendInvite from "./FriendInvite";

const FriendInviteList = () => {
    return (
        <div
            id="friendInviteList"
            className="w-full h-full px-2 py-5 flex-col relative bg-white/25 rounded-[30px] gap-2"
        >
            <FriendInvite />
            <FriendInvite />
            <FriendInvite />
        </div>
    );
};

export default FriendInviteList;
