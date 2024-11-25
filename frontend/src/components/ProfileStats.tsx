"use client";
import React, { useEffect, useState } from "react";

const ProfileStats = ({ userId }: { userId: string }) => {
    const defaultStats = {
        id: userId,
        numFollowers: 0,
        numFollowing: 0,
        numPosts: 0,
        numLikes: 0,
    } as const;

    const [stats, setStats] = useState<typeof defaultStats>(defaultStats);

    useEffect(() => {
        fetch(`/api/user/${userId}/stats`)
            .then((resp) => {
                if (resp.ok) return resp.json();
                throw "error";
            })
            .then(setStats)
            .catch(console.error);
    }, [userId]);

    return (
        <div className="w-[400px] h-16 bg-white/30 rounded-2xl flex flex-row items-center justify-between px-3">
            <div className="flex flex-col items-center w-[86px]">
                <div className="font-bold">Posts</div>
                <div className="text-black/50">{stats.numPosts}</div>
            </div>
            <div className="flex flex-col items-center w-[86px]">
                <div className="font-bold">Likes</div>
                <div className="text-black/50">{stats.numLikes}</div>
            </div>
            <div className="flex flex-col items-center w-[86px]">
                <div className="font-bold">Follow(s)</div>
                <div className="text-black/50">{stats.numFollowing}</div>
            </div>
            <div className="flex flex-col items-center w-[86px]">
                <div className="font-bold">Follower(s)</div>
                <div className="text-black/50">{stats.numFollowers}</div>
            </div>
        </div>
    );
};

export default ProfileStats;
