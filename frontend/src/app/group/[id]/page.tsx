"use client";

import React, { useEffect, useState } from "react";
import { Group } from "@/types/group";
import NewEvent from "./NewEvent";
import Events from "./Events";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { BackIcon } from "@/components/icons/BackIcon";
import Link from "next/link";
import Image from "next/image";
import { NewPost } from "@/components/NewPost";
import { Post } from "@/types/post";
import PostComponent from "@/components/PostComponent";
import { useParams } from "next/navigation";
import { MemberGroupList } from "./MemberGroupList";
import { User } from "@/types/user";
import { FollowersList } from "@/components/FollowingList";

export default function GroupPage() {
    const { id } = useParams<{ id: string }>();
    const [group, setGroup] = useState<Group | null>(null);
    const [loading, setLoading] = useState(true);
    const [posts, setPosts] = useState<Post[] | null>(null);

    const [showMemberList, setShowMemberList] = useState(false)
    const [showEventList, setShowEventList] = useState(true)


    useEffect(() => {
        const fetchInfos = async () => {
            try {
                const response = await fetch(`/api/groups/${id}`);
                const group: Group = await response.json();
                setGroup(group);
                const test = await fetch(
                    `/api/groups/${id}/posts?limit=20&offset=0`
                );
                if (!test.ok) throw "Error fetching posts";
                const posts: Post[] = await test.json();
                setPosts(posts);
            } catch (error) {
                console.error(error);
            } finally {
                setLoading(false);
            }
        };

        fetchInfos();
    }, [id]);

    if (loading) return <>loading</>;

    if (!group) return <>Group Not Found</>;
    // const { limit, offset, next, previous } = useQueryParams();

    const handleMemberListClick = () => {
        setShowMemberList(true)
        setShowEventList(false)
    }

    const handleEventListClick = () => {
        setShowMemberList(false)
        setShowEventList(true)
    }

    const handleRequestClick = () => {
        //Request handler
        console.log("Sending request to join");
        fetch(`/api/groups/${group.id}/request`, { method: "POST" });
        console.log("Request Send !");
    }

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <div className="flex flex-row justify-between items-center w-full h-16 px-5 py-2 shadow-xl">
                        <Link href={"/group"}>
                            <BackIcon />
                        </Link>

                        <div className="flex flex-col justify-center items-center">
                            <h1 className="font-bold">{group.name}</h1>
                            <p>{group.description}</p>
                        </div>

                        <div className="flex flex-row items-center gap-5 text-3xl">
                            <Image
                                src={group.image}
                                alt=""
                                width={50}
                                height={50}
                            />

                            {/* displaying only if in group */}
                            {/* Invite followers to group */}
                            <Link href={`/group/${group.id}/chatroom`}>Chat</Link>

                            <input
                                type="button"
                                value="+"
                                className="font-bold"
                            />
                        </div>
                    </div>

                    {/* Display if not in group */}
                    {/* <div className="flex flex-col items-center font-bold text-3xl gap-5">
                        <h2>You are not in the group yet, <br /> click below to send a request !</h2>

                        <label htmlFor="request-to-group"></label>
                        <input name="request-to-group" id="request-to-group" type="button" value="request to join" onClick={handleRequestClick} />
                    </div> */}

                    {/* Display if in group */}
                    <div className="flex flex-row w-full h-full">
                        <div className="flex flex-col items-center w-72 border-r-4">
                            <div className="flex flex-col pt-3 gap-2">
                                <NewPost groupId={id} />
                                <NewEvent groupId={group.id} />
                            </div>

                            <ul className="flex flex-col items-center">
                                <li onClick={handleMemberListClick} className="font-bold cursor-pointer">Members</li>
                                {showMemberList && <MemberGroupList />}


                                <li onClick={handleEventListClick} className="font-bold cursor-pointer">Events</li>
                                {showEventList && <Events groupId={group.id} />}
                            </ul>
                        </div>
                        {posts && (
                            <div className="flex flex-col w-full p-3 gap-3 overflow-scroll no-scrollbar">
                                {posts.map((post, idx) => (
                                    <PostComponent key={idx} post={post} />
                                ))}
                            </div>
                        )}
                    </div>

                    <span className="absolute top-0 right-0 translate-x-full translate-y-40"><FollowersList groupId={group.id} /></span>
                </div>
            </HomeProfileLayout>
        </>
    );
}
