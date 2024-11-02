"use client";

import React from "react";

export default function FollowButton({
    userId,
    username,
}: {
    userId: string;
    username: string;
}) {
    const follow = () =>
        fetch(`/api/user/${userId}/follow`, { method: "POST" });
    return (
        <button className="bg-gray-300 p-2" onClick={follow}>
            Follow {username}
        </button>
    );
}