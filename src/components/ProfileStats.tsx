"use client";
import React, { useEffect, useState } from "react";

type UserStats = {
    id: string;
    numFollowers: number;
    numFollowing: number;
    numPosts: number;
    numLikes: number;
};

const defaultStats = {
    id: "",
    numFollowers: 0,
    numFollowing: 0,
    numPosts: 0,
    numLikes: 0,
};

const ProfileStats = ({ userId }: { userId: string }) => {
    const [stats, setStats] = useState<UserStats>(defaultStats);

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
                <div className="font-bold">Following</div>
                <div className="text-black/50">{stats.numFollowing}</div>
            </div>
            <div className="flex flex-col items-center w-[86px]">
                <div className="font-bold">Followers</div>
                <div className="text-black/50">{stats.numFollowers}</div>
            </div>
        </div>
    );
};

export default ProfileStats;
