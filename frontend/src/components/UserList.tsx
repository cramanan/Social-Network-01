"use client";

import React from "react";
import Users from "./Users";
export default function UserList() {
    return (
        <>
            <div
                id="userList"
                className="flex flex-col w-full h-[calc(100vh-111px)] xl:w-60 xl:h-fit xl:bg-white/40 xl:rounded-3xl xl:py-3"
            >
                <h2 className="text-4xl text-white text-center py-5 xl:sr-only">Friend List</h2>

                <div className="flex flex-col items-center gap-3 mx-5 overflow-scroll no-scrollbar xl:h-[75vh]">
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
                    <Users />
                    <Users />
                </div>
            </div>
        </>
    );
}
