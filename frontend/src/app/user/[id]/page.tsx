"use client";

import FollowButton from "@/components/FollowButton";
import PostComponent from "@/components/PostComponent";
import { PostMedia } from "@/components/PostMedia";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { Post } from "@/types/post";
import { User } from "@/types/user";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function Page() {
    const [user, setUser] = useState<User | null>();
    const { id } = useParams<{ id: string }>();
    const [posts, setPosts] = useState<Post[]>([]);
    const [isNotFollowing, setIsNotFollowing] = useState(false)

    useEffect(() => {
        const fetchUser = async () => {
            const response = await fetch(`/api/users/${id}`);
            setIsNotFollowing(() => {
                return response.status === 401
            })

            const user = await response.json();
            setUser(user);
        };
        fetchUser();
    }, [id]);

    useEffect(() => {
        const fetchPosts = async () => {
            const response = await fetch("/api/profile/posts"); // TODO: change route to users posts
            const data: Post[] = await response.json();
            setPosts(data);
        };

        fetchPosts();
    }, []);

    if (!user) return <></>;

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center">
                    <div className="flex flex-col justify-center items-end my-3 mt-11">
                        <ProfileBanner {...user} />
                        <ProfileStats userId={user.id} />
                    </div>

                    <FollowButton {...user} />

                    {!user.isPrivate || !isNotFollowing ?
                        (<div className="flex flex-col h-[calc(100vh-360px)] overflow-scroll no-scrollbar gap-2 pb-2 xl:w-[1000px] xl:h-[calc(100vh-300px)]">
                            {posts.length > 0 ? (
                                posts.map((post, idx) =>
                                    post.images.length > 0 ? (
                                        <PostMedia key={idx} post={post} />
                                    ) : (
                                        <PostComponent key={idx} post={post} />
                                    )
                                )
                            ) : (
                                <p className="text-3xl font-bold text-center mt-10">No posts found</p>
                            )}
                        </div>)
                        :
                        (<h2 className="font-bold text-5xl">Account is in private</h2>)
                    }
                </div>
            </HomeProfileLayout>
        </>
    );
}
