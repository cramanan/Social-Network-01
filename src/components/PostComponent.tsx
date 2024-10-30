"use client";

import React, { useState } from "react";
import Comment from "./Comment";
import { RedLikeIcon } from "./icons/RedLikeIcon";
import { BookmarkIcon } from "./icons/BookmarkIcon";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import { Post } from "@/types/post";
import Image from "next/image";
import Link from "next/link";

const PostComponent = ({ post }: { post: Post }) => {
    const [isLiked, setIsLiked] = useState(false);

    const handleLikeClick = () => setIsLiked(!isLiked);

    return (
        <>
            <div className="flex flex-col relative w-full bg-white/95 rounded-[30px]">
                <div className="flex flex-row justify-between items-center mr-5">
                    <div className="flex flex-row items-center ml-2 mt-2 gap-5">
                        <div className="flex flex-col">
                            <div className="w-12 h-12 bg-[#af5f5f] rounded-[100px]"></div>
                            <Link
                                href={`/user/${post.userId}`}
                                className="h-[21px] text-black text-xl font-semibold font-['Inter']"
                            >
                                {post.username}
                            </Link>
                            <span className="h-[29px] text-black/50 text-base font-extralight font-['Inter']">
                                Friday 6 september 16:03
                            </span>
                        </div>
                    </div>
                    <BookmarkIcon />
                </div>
                <div className="h-[110px] line-clamp-5 overflow-hidden text-black text-base font-normal font-['Inter'] leading-[22px] m-5 mr-7 mb-10">
                    {post.images.map((src, idx) => (
                        <a href={src} key={idx} target="_blank">
                            <Image src={src} width={100} height={100} alt="" />
                        </a>
                    ))}
                </div>
                <div className="h-[110px] line-clamp-5 overflow-hidden text-black text-base font-normal font-['Inter'] leading-[22px] m-5 mr-7 mb-10">
                    {post.content}
                </div>
                <div className="flex flex-row gap-5 ml-5">
                    <div onClick={handleLikeClick}>
                        {isLiked ? <RedLikeIcon /> : <LikeIcon />}
                    </div>
                    <CommentIcon />
                </div>
                <div className="mb-5 mt-1 ml-5 mr-10">
                    <Comment />
                    <Comment />
                </div>
                <div className="text-center text-black text-sm font-medium font-['Inter'] mb-2 cursor-pointer">
                    See more
                </div>
            </div>
        </>
    );
};

export default PostComponent;
