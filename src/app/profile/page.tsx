"use client";

import Actualite from "@/components/Actualite";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useAuth } from "@/providers/AuthContext";
import { redirect } from "next/navigation";

export default function Profile() {
	const { user } = useAuth();

	if (!user) redirect("/auth");

	return (
		<HomeProfileLayout>
			<ProfileBanner nickname={user.nickname} firstName={user.firstName} />
			<ProfileStats />
			<Actualite />
		</HomeProfileLayout>
	);
}
