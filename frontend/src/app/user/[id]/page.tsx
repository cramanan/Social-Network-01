"use client";

import FollowButton from "@/components/FollowButton";
import ProfileStats from "@/components/ProfileStats";
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
            <div className="whitespace-pre-wrap">
                {JSON.stringify(user, null, "\t")}
            </div>
            <ProfileStats userId={user.id} />
            <FollowButton userId={user.id} username={user.nickname} />
        </>
    );
}
