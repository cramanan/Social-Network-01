"use client";

import React, { useEffect, useState, ChangeEvent } from "react";
import { Group } from "@/types/group";
import { ServerChat } from "@/types/chat";
import NewEvent from "./NewEvent";
import Events from "./Events";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { BackIcon } from "@/components/icons/BackIcon";
import { SendIcon } from "@/components/icons/SendIcon";
import { EmoteIcon } from "@/components/icons/EmoteIcon";
import Link from "next/link";
import Image from "next/image";
import { NewPost } from "@/components/NewPost";
import { Post } from "@/types/post";
import PostComponent from "@/components/PostComponent";
import { MemberGroupList } from "./MemberGroupList";
import { FollowersList } from "@/components/FollowingList";
import { useParams } from "next/navigation";
import useQueryParams from "@/hooks/useQueryParams";
import { useAuth } from "@/hooks/useAuth";
import formatDate from "@/utils/formatDate";
import EmojiPicker, { EmojiClickData } from "emoji-picker-react";

export default function GroupPage() {
    const { user, loading: authLoading } = useAuth();
    const { id } = useParams<{ id: string }>();
    const [group, setGroup] = useState<Group | null>(null);
    const [loading, setLoading] = useState(true);
    const [posts, setPosts] = useState<Post[] | null>(null);
    const [unauthorized, setUnauthorized] = useState(false);

    // Chat states
    const [showChat, setShowChat] = useState(false);
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const [messages, setMessages] = useState<ServerChat[]>([]);
    const [content, setContent] = useState("");
    const [showEmojiPicker, setShowEmojiPicker] = useState(false);
    const { limit, offset } = useQueryParams();

    const [showMemberList, setShowMemberList] = useState(false);
    const [showEventList, setShowEventList] = useState(true);

    const [showAddPeople, setShowAddPeople] = useState(false);
    const handleAddPeople = () => setShowAddPeople(!showAddPeople);

    useEffect(() => {
        const fetchInfos = async () => {
            try {
                const response = await fetch(`/api/groups/${id}`);
                const group: Group = await response.json();
                setUnauthorized(() => {
                    return response.status === 401;
                });
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

    // Chat useEffects
    useEffect(() => {
        if (!showChat || unauthorized) return;

        const fetchMessages = async () => {
            try {
                const response = await fetch(
                    `/api/groups/${id}/chats?limit=${limit}&offset=${offset}`
                );
                const data = await response.json();
                setMessages(data);
            } catch (error) {
                console.error(error);
            }
        };

        fetchMessages();
    }, [id, limit, offset, unauthorized, showChat]);

    useEffect(() => {
        if (!showChat && unauthorized) return;

        const ws = new WebSocket(
            `ws://${process.env.NEXT_PUBLIC_API_URL}/api/groups/${id}/chatroom`
        );

        ws.addEventListener("message", (msg) => {
            const message = JSON.parse(msg.data) as ServerChat;
            setMessages((prev) => [...prev, message]);
        });

        setSocket(ws);

        return () => {
            ws.close();
            setSocket(null);
        };
    }, [id, unauthorized, showChat]);

    const handleEmojiClick = (emojiData: EmojiClickData) => {
        setContent((prev) => prev + emojiData.emoji);
        setShowEmojiPicker(false);
    };

    const onSubmit = async (e: ChangeEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!user || !content.trim()) return;

        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify({ content }));
            setMessages((prev) => [
                ...prev,
                {
                    senderId: user.id,
                    content,
                    timestamp: formatDate(new Date().toString()),
                    recipientId: id,
                },
            ]);
            setContent("");
        }
    };

    const handleMemberListClick = () => {
        setShowMemberList(true);
        setShowEventList(false);
    };

    const handleEventListClick = () => {
        setShowMemberList(false);
        setShowEventList(true);
    };

    const handleRequestClick = () => {
        console.log("Sending request to join");
        fetch(`/api/groups/${group?.id}/request`, { method: "POST" });
        console.log("Request Send !");
    };

    const toggleChat = () => {
        setShowChat((prev) => !prev);
    };

    if (loading || authLoading) return <>loading</>;
    if (!group) return <>Group Not Found</>;
    if (!user)
        return (
            <div className="flex items-center justify-center h-screen">
                Please log in to continue
            </div>
        );

    return (
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

                        {!unauthorized && (
                            <>
                                <button onClick={toggleChat}>Chat</button>
                                <input
                                    onClick={handleAddPeople}
                                    type="button"
                                    value="+"
                                    className="font-bold"
                                />
                            </>
                        )}
                    </div>
                </div>

                {unauthorized ? (
                    <div className="flex flex-col items-center font-bold text-3xl gap-5">
                        <h2>
                            You are not in the group yet, <br /> click below to
                            send a request !
                        </h2>
                        <label htmlFor="request-to-group"></label>
                        <input
                            name="request-to-group"
                            id="request-to-group"
                            type="button"
                            value="request to join"
                            onClick={handleRequestClick}
                        />
                    </div>
                ) : (
                    <>
                        <div className="flex flex-row w-full h-full">
                            <div className="flex flex-col items-center w-72 border-r-4">
                                <div className="flex flex-col pt-3 gap-2">
                                    <NewPost groupId={id} />
                                    <NewEvent groupId={group.id} />
                                </div>

                                <ul className="flex flex-col items-center">
                                    <li
                                        onClick={handleMemberListClick}
                                        className="font-bold cursor-pointer"
                                    >
                                        Members
                                    </li>
                                    {showMemberList && (
                                        <MemberGroupList groupId={group.id} />
                                    )}

                                    <li
                                        onClick={handleEventListClick}
                                        className="font-bold cursor-pointer"
                                    >
                                        Events
                                    </li>
                                    {showEventList && (
                                        <Events groupId={group.id} />
                                    )}
                                </ul>
                            </div>

                            <div className="flex flex-col w-full h-full">
                                {showChat ? (
                                    <>
                                        <ul className="flex flex-col flex-grow px-3 py-2 overflow-scroll no-scrollbar">
                                            {messages.map(
                                                (
                                                    {
                                                        senderId,
                                                        content,
                                                        timestamp,
                                                    },
                                                    idx
                                                ) => {
                                                    const isCurrentUser =
                                                        senderId === user.id;
                                                    return (
                                                        <li
                                                            key={idx}
                                                            className={`flex flex-col ${
                                                                isCurrentUser
                                                                    ? "self-end items-end"
                                                                    : "self-start"
                                                            } mb-3`}
                                                        >
                                                            <p
                                                                className={`p-3 rounded-2xl w-fit max-w-[80%] break-words ${
                                                                    isCurrentUser
                                                                        ? "bg-[#b88ee5] text-black"
                                                                        : "bg-[#4174e2] text-white"
                                                                }`}
                                                            >
                                                                {content}
                                                            </p>
                                                            <div className="text-sm text-gray-500 mt-1">
                                                                {timestamp}
                                                            </div>
                                                        </li>
                                                    );
                                                }
                                            )}
                                        </ul>

                                        {showEmojiPicker && (
                                            <div className="absolute bottom-20 left-72">
                                                <EmojiPicker
                                                    onEmojiClick={
                                                        handleEmojiClick
                                                    }
                                                />
                                            </div>
                                        )}

                                        <form
                                            onSubmit={onSubmit}
                                            className="h-[50px] flex flex-row items-center m-5 bg-[#445ab3]/20 rounded-[25px] p-2 gap-2"
                                        >
                                            <button
                                                type="button"
                                                onClick={() =>
                                                    setShowEmojiPicker(
                                                        !showEmojiPicker
                                                    )
                                                }
                                                className="p-2"
                                            >
                                                <EmoteIcon />
                                            </button>
                                            <input
                                                type="text"
                                                placeholder="Enter your message"
                                                value={content}
                                                onChange={(e) =>
                                                    setContent(e.target.value)
                                                }
                                                className="bg-transparent w-full placeholder:text-black outline-none"
                                            />
                                            <button
                                                type="submit"
                                                className="p-2"
                                            >
                                                <SendIcon />
                                            </button>
                                        </form>
                                    </>
                                ) : (
                                    <div className="flex flex-col w-full p-3 gap-3 overflow-scroll no-scrollbar xl:h-[calc(100vh-135px)]">
                                        {posts &&
                                            posts.map((post, idx) => (
                                                <PostComponent
                                                    key={idx}
                                                    post={post}
                                                />
                                            ))}
                                    </div>
                                )}
                            </div>
                        </div>

                        {showAddPeople && (
                            <span className="absolute top-0 right-0 translate-x-full translate-y-40">
                                <FollowersList groupId={group.id} />
                            </span>
                        )}
                    </>
                )}
            </div>
        </HomeProfileLayout>
    );
}
