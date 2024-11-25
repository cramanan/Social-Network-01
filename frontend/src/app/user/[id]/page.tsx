"use client";

import FollowButton from "@/components/FollowButton";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { Params } from "@/types/query";
import { User } from "@/types/user";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function Page() {
    const [user, setUser] = useState<User | null>();
    const { id } = useParams();
    console.log(id);

    useEffect(() => {
        const fetchUser = async () => {
            const response = await fetch(`/api/user/${id}`);
            const user = await response.json();
            setUser(user);
        };
        fetchUser();
    }, []);
    if (!user) return <></>;

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col justify-center items-end my-3 mt-11">
                    <ProfileBanner {...user} />
                    <ProfileStats userId={user.id} />
                </div>
                <FollowButton userId={user.id} username={user.nickname} />
                <div className="whitespace-pre-wrap">
                    {JSON.stringify(user, null, "\t")}
                </div>
            </HomeProfileLayout>
        </>
    );
}
