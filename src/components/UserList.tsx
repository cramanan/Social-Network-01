"use client";

import React from "react";
import Users from "./Users";
export default function UserList() {
    return (
        <>
            <div
                id="userList"
                className="w-60 h-fit bg-white bg-opacity-40 rounded-3xl py-3"
            >
                <div className="flex flex-col gap-3 mx-5">
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
