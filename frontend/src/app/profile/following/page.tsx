"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { User } from "@/types/user";
import Image from "next/image";
import React, { useEffect, useState } from "react";

export default function Page() {
    const [following, setFollowing] = useState<User[]>([]);
    const { limit, offset, next, previous } = useQueryParams();
    useEffect(() => {
        const fetchFollowing = async () => {
            const response = await fetch(
                `/api/profile/following?limit=${limit}&offset=${offset}`
            );
            const data: User[] = await response.json();
            setFollowing(data);
        };

        fetchFollowing();
    }, [limit, offset]);
    return (
        <div>
            {following.map(({ id, nickname, image }, idx) => (
                <a href={`/user/${id}`} key={idx} className="flex items-center">
                    <Image
                        alt=""
                        width={100}
                        height={100}
                        src={image}
                        className="h-20 w-20 rounded-full object-cover"
                    />
                    <span>{nickname}</span>
                </a>
            ))}
            <button onClick={next}>Next</button>
            <button onClick={previous}>Previous</button>
        </div>
    );
}
