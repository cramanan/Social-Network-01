"use client";

import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import { useAuth } from "@/hooks/useAuth";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { redirect } from "next/navigation";

export default function Profile() {
    const { user, loading } = useAuth();

    if (loading) return <>Loading</>;

    if (!user) redirect("/auth");

    return (
        <HomeProfileLayout>
            <div className="flex flex-col items-center gap-2">
                <ProfileBanner
                    id={user.id}
                    firstName={user.firstName}
                    image={user.image}
                />
                <ProfileStats userId={user.id} />
            </div>
        </HomeProfileLayout>
    );
}
