"use client";

import { BackIcon } from "@/components/icons/BackIcon";
import useQueryParams from "@/hooks/useQueryParams";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { User } from "@/types/user";
import Image from "next/image";
import React, { useEffect, useState } from "react";

export default function Page() {
    const [followers, setFollowers] = useState<User[]>([]);
    const { limit, offset } = useQueryParams();
    useEffect(() => {
        const fetchFollowers = async () => {
            const response = await fetch(
                `/api/profile/followers?limit=${limit}&offset=${offset}`
            );
            const data: User[] = await response.json();
            setFollowers(data);
        };

        fetchFollowers();
    }, [limit, offset]);
    return (
        <>
            <HomeProfileLayout >
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <div className="shadow-xl w-full mb-5 p-3">
                        <div className="flex justify-between">
                            <a href="/profile"><BackIcon /></a>
                            <h2 className="text-black text-xl font-bold font-['Inter'] tracking-wide text-center">Follower(s)</h2>
                            <span></span>
                        </div>
                    </div>

                    {followers.length > 0 ?
                        followers.map(({ id, nickname, image }, idx) => (
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
                        ))
                        :
                        (
                            <p className="text-center font-bold">
                                No follower(s) found.
                            </p>
                        )}

                </div>
            </HomeProfileLayout>
        </>
    );
}
