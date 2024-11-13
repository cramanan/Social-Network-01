"use client";

import ProfileBanner from "@/components/ProfileBanner";
import ProfilePost from "@/components/ProfilePost";
import ProfileStats from "@/components/ProfileStats";
import { useAuth } from "@/hooks/useAuth";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { Post } from "@/types/post";
import { redirect } from "next/navigation";
import { useEffect, useState } from "react";

export default function Profile() {
    const { user, loading } = useAuth();
    const [posts, setPosts] = useState<Post[]>([]);

    useEffect(() => {
        const fetchPosts = async () => {
            const response = await fetch("/api/profile/posts");
            const data: Post[] = await response.json();
            setPosts(data);
        };

        fetchPosts();
    }, []);

    if (loading) return <>Loading</>;

    if (!user) redirect("/auth");

    return (
        <HomeProfileLayout>
            <div className="flex flex-col justify-center items-center">
                <div className="flex flex-col justify-center items-center my-3">
                    <ProfileBanner {...user} />
                    <ProfileStats userId={user.id} />
                </div>

                <div className="flex flex-col gap-2 h-[calc(100vh-300px)] overflow-scroll no-scrollbar">
                    {posts.map((post, idx) => (
                        <ProfilePost key={idx} {...post} />
                    ))}
                </div>
            </div>
        </HomeProfileLayout>
    );
}
