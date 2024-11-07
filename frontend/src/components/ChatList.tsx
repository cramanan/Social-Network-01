import React from 'react'
import Users from "./Users"

const ChatList = () => {
    return (
        <>

            <div className="relative flex flex-col w-full h-[calc(100vh-111px)] xl:w-60 xl:h-fit xl:rounded-3xl xl:py-3 xl:bg-white/40" aria-label="Chat list" role="region">
                <h2 className="text-4xl text-white text-center py-5 xl:sr-only">Chat List</h2>

                <div className="flex flex-col gap-3 mx-5 overflow-scroll no-scrollbar xl:max-h-[75vh] xl:h-fit">
                    {/* <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users /> */}
                </div>
            </div>
        </>
    )
}

export default ChatList