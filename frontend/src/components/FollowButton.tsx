"use client";

import React from "react";

export default function FollowButton({
    id,
    nickname,
}: {
    id: string;
    nickname: string;
}) {
    const follow = () =>
        fetch(`/api/user/${id}/send-request`, { method: "POST" });
    return (
        <button className="bg-gray-300 w-fit rounded-xl p-2" onClick={follow}>
            Follow {nickname}
        </button>
    );
}
