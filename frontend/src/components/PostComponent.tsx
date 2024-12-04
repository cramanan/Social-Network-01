"use client";

import React, {
    useEffect,
    useRef,
    useState,
} from "react";
import { CommentIcon } from "./icons/CommentIcon";
import { Comment as CommentType, Post } from "@/types/post";
import Image from "next/image";
import Link from "next/link";
import formatDate from "@/utils/formatDate";
import { LikeIcon } from "./icons/LikeIcon";
import Comment from "./Comment";
import { NewComment } from "./NewComment";
import { useAuth } from "@/hooks/useAuth";

const PostComponent = ({ post }: { post: Post }) => {
    const { user } = useAuth()
    const [isExpanded, setIsExpanded] = useState(false);
    const [ShowAllComment, setShowAllComment] = useState(false);
    const [isOverflowing, setIsOverflowing] = useState(false);
    const contentRef = useRef<HTMLDivElement | null>(null);
    const handleSeeMore = () => setIsExpanded(!isExpanded);
    const handleShowAllComment = () => setShowAllComment(!ShowAllComment);

    const [allComments, setAllComments] = useState<CommentType[]>([]);

    useEffect(() => {
        const fetchComments = async () => {
            const response = await fetch(`/api/posts/${post.id}/comments`);
            const data: CommentType[] = await response.json();

            setAllComments(data);
        };

        fetchComments();
    }, [post.id]);

    useEffect(() => {
        const checkOverflow = () => {
            if (contentRef.current) {
                // Check if content is longer than the container height
                const isOverflowing =
                    contentRef.current.scrollHeight >
                    contentRef.current.clientHeight;
                setIsOverflowing(isOverflowing);
            }
        };

        checkOverflow();
        window.addEventListener("resize", checkOverflow);

        return () => window.removeEventListener("resize", checkOverflow);
    }, [post.content]);

    return (
        <>
            <div className="flex flex-col relative w-full bg-white/95 xl:rounded-[30px]">
                <div className="flex flex-row justify-between items-center pr-5 mb-3">
                    <div className="flex flex-row items-center ml-2 mt-2 gap-3">
                        <Link href={`${user?.id === post.userId ? `/profile` : `/user/${post.userId}`}`}>
                            <Image
                                src={post.userImage}
                                width={48}
                                height={48}
                                alt=""
                                className="flex justify-center items-center w-12 h-12 border border-black rounded-full"
                            />
                        </Link>

                        <div className="flex flex-col">
                            <Link
                                href={`${user?.id === post.userId ? `/profile` : `/user/${post.userId}`}`}
                                className="text-black text-xl font-semibold font-['Inter']"
                            >
                                {post.username}
                            </Link>
                            <span className="text-black/50 text-base font-extralight font-['Inter']">
                                {formatDate(post.timestamp)}
                            </span>
                        </div>
                    </div>
                </div>

                {post.images.length > 0 && (
                    <p className="w-fit mx-5 mb-0">
                        {post.images.map((src, idx) => (
                            <a
                                href={src}
                                key={idx}
                                target="_blank"
                                rel="noopener noreferrer"
                            >
                                <Image
                                    src={src}
                                    width={100}
                                    height={100}
                                    alt=""
                                    className="max-h-[100px] w-auto h-auto object-contain"
                                />
                            </a>
                        ))}
                    </p>
                )}

                <p
                    ref={contentRef}
                    className={`h-fit text-black text-base font-normal font-['Inter'] leading-[22px] text-justify m-5 mr-10 whitespace-pre-wrap ${isExpanded
                        ? ""
                        : "h-[110px] line-clamp-5 overflow-hidden"
                        }`}
                >
                    {post.content}
                </p>

                {isOverflowing && (
                    <button
                        onClick={handleSeeMore}
                        className="text-center text-black text-sm font-medium font-['Inter']"
                    >
                        {isExpanded ? "See less" : "See more"}
                    </button>
                )}

                <div className="flex flex-row gap-10 ml-5">
                    <button className="flex gap-2">
                        <LikeIcon />
                        0
                    </button>
                    <div className="flex gap-2">
                        <CommentIcon />
                        {allComments.length}
                    </div>
                </div>

                <div
                    className={`bg-black/10 overflow-hidden my-3 ml-5 mr-10 ${ShowAllComment ? "h-fit" : "max-h-[108px]"
                        }`}
                >
                    {allComments.map((comment, idx) => (
                        <Comment key={idx} {...comment} />
                    ))}
                </div>

                {allComments.length > 2 && (
                    <div className="text-center text-black text-sm font-medium font-['Inter'] mb-2">
                        <button
                            onClick={handleShowAllComment}
                            className="cursor-pointer"
                        >
                            {ShowAllComment ? "Less comments" : "More comments"}
                        </button>
                    </div>
                )}

                <NewComment {...post} />
            </div>
        </>
    );
};

export default PostComponent;
