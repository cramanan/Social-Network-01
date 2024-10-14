'use client'

import React from "react";
import Users from "./Users";
export default function UserList() {
    return (
        <>
            <div id="userList" className="flex w-44 h-fit bg-white flex-col items-start gap-3 bg-opacity-40 rounded-3xl pt-3 pb-4">
                <Users />
                <Users />
                <Users />
                <Users />
                <Users />
                <Users />
                <Users />
            </div>
        </>
    );
}