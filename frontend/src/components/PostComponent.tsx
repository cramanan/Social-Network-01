"use client";

import React, { useEffect, useRef, useState } from "react";
import Comment from "./Comment";
import { BookmarkIcon } from "./icons/BookmarkIcon";
import { CommentIcon } from "./icons/CommentIcon";
import { Post } from "@/types/post";
import Image from "next/image";
import Link from "next/link";
import formatDate from "@/utils/formatDate";
import { LikeIcon } from "./icons/LikeIcon";

const PostComponent = ({ post }: { post: Post }) => {
    const [isLiked, setIsLiked] = useState(false);
    const [isExpanded, setIsExpanded] = useState(false);
    const [ShowAllComment, setShowAllComment] = useState(false);
    const [isOverflowing, setIsOverflowing] = useState(false);
    const contentRef = useRef<HTMLAnchorElement>(null);

    const handleLikeClick = () => setIsLiked(!isLiked);
    const handleSeeMore = () => setIsExpanded(!isExpanded);
    const handleShowAllComment = () => setShowAllComment(!ShowAllComment);

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
                        <Link href={`/user/${post.userId}`}>
                            <Image
                                src={"/"}
                                width={48}
                                height={48}
                                alt=""
                                className="flex justify-center items-center w-12 h-12 border border-black rounded-full"
                            ></Image>
                        </Link>

                        <div className="flex flex-col">
                            <Link
                                href={`/user/${post.userId}`}
                                className="text-black text-xl font-semibold font-['Inter']"
                            >
                                {post.username}
                            </Link>
                            <span className="text-black/50 text-base font-extralight font-['Inter']">
                                {formatDate(post.timestamp)}
                            </span>
                        </div>
                    </div>
                    <BookmarkIcon />
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

                <Link
                    ref={contentRef}
                    href={`/post/${post.id}`}
                    className={`h-fit text-black text-base font-normal font-['Inter'] leading-[22px] m-5 mr-10 ${isExpanded
                        ? ""
                        : "h-[110px] line-clamp-5 overflow-hidden"
                        }`}
                >
                    {post.content}
                </Link>

                {isOverflowing && (
                    <button
                        onClick={handleSeeMore}
                        className="text-center text-black text-sm font-medium font-['Inter']"
                    >
                        {isExpanded ? "See less" : "See more"}
                    </button>
                )}

                <div className="flex flex-row gap-10 ml-5">
                    <button onClick={handleLikeClick}>
                        <LikeIcon isLiked={isLiked} />
                    </button>
                    <CommentIcon />
                </div>

                <div
                    className={`bg-black/10 overflow-hidden my-3 ml-5 mr-10 ${ShowAllComment ? "h-fit" : "max-h-[108px]"
                        }`}
                >
                    <Comment />
                    <Comment />
                </div>

                <div className="text-center text-black text-sm font-medium font-['Inter'] mb-2">
                    <button
                        onClick={handleShowAllComment}
                        className="cursor-pointer"
                    >
                        {ShowAllComment ? "Less comments" : "More comments"}
                    </button>
                </div>

                <div className="h-[58px] pl-px pr-3 pt-[11px] pb-[7px] bg-[#f2eeee] rounded-[10px] gap-2 items-center inline-flex mx-5 my-2">
                    <div className="w-full flex flex-row items-center gap-2">
                        <div className="w-[44px] h-[40px] relative">emote</div>
                        <input
                            type="text"
                            placeholder="Enter your comment"
                            className="w-full h-[30px] text-black text-xl font-extralight font-['Inter'] bg-white/0"
                        ></input>
                    </div>
                    <div className="self-stretch pl-[11px] pr-3 pt-[5px] bg-gradient-to-t from-[#e1d3eb] via-[#6f46c0] to-[#e0d3ea] rounded-[30px] justify-center items-center inline-flex">
                        <button className="h-[25px] text-center text-black text-[15px] font-medium font-['Inter']">
                            Send
                        </button>
                    </div>
                </div>
            </div>
        </>
    );
};

export default PostComponent;
