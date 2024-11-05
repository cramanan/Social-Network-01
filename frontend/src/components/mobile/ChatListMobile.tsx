import React from 'react'
import Users from "../Users"

export const ChatListMobile = () => {
    return (
        <>
            <div className="relative z-5 flex flex-col w-full">
                <h2 className="text-4xl text-white text-center py-5">Mobile Chat</h2>

                <div className="flex flex-col h-[75vh] gap-3 mx-5 overflow-scroll">
                    <Users />
                    <Users />
                    <Users />
                </div>
            </div>
        </>
    )
}
