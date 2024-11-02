"use client";

import Actualite from "@/components/Actualite";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useAuth } from "@/providers/AuthContext";
import { redirect } from "next/navigation";

export default function Profile() {
    const { user, loading } = useAuth();

    if (loading) return <>Loading</>;

    if (!user) redirect("/auth");

    return (
        <HomeProfileLayout>
            <div className="flex flex-col items-center">
                <ProfileBanner
                    id={user.id}
                    firstName={user.firstName}
                    image={user.image}
                />
                <ProfileStats userId={user.id} />
                <div>
                    <Actualite />
                </div>
            </div>
        </HomeProfileLayout>
    );
}
