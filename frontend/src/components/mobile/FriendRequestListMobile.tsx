import React from 'react'
import FriendInvite from "../FriendInvite"

export const FriendRequestListMobile = () => {
    return (
        <>
            <div className="relative z-5 flex flex-col w-full">
                <h2 className="text-4xl text-white text-center py-5">Friend Request List</h2>

                <div className="flex flex-col h-[70vh] items-center gap-3 mx-5 overflow-scroll">
                    <FriendInvite />
                    <FriendInvite />
                    <FriendInvite />
                    <FriendInvite />
                    <FriendInvite />
                    <FriendInvite />
                    <FriendInvite />
                </div>
            </div>
        </>
    )
}
