"use client";

import PostComponent from "@/components/PostComponent";
import { PostMedia } from "@/components/PostMedia";
import ProfileBanner from "@/components/ProfileBanner";
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
            <div className="h-[calc(100vh-185px)] flex flex-col justify-center items-center xl:h-[calc(100vh-60px)]">
                <div className="flex flex-col justify-center items-end my-3 mt-11">
                    <ProfileBanner {...user} />
                    <ProfileStats userId={user.id} />
                </div>

                <div className="flex flex-col h-[calc(100vh-260px)] overflow-scroll no-scrollbar gap-2 pb-2 xl:w-[1000px]">
                    {posts.map((post, idx) => (
                        post.images.length > 0 ? (
                            <PostMedia key={idx} post={post} />
                        ) : (
                            <PostComponent key={idx} post={post} />
                        )
                    ))}
                </div>
            </div>
        </HomeProfileLayout>
    );
}
