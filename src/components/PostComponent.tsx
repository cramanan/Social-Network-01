"use client";

import React, { useState } from "react";
import Comment from "./Comment";
import { BookmarkIcon } from "./icons/BookmarkIcon";
import { CommentIcon } from "./icons/CommentIcon";
import { Post } from "@/types/post";
import Image from "next/image";
import Link from "next/link";
import formatDate from "@/utils/formatDate";

const PostComponent = ({ post }: { post: Post }) => {
    const [isLiked, setIsLiked] = useState(false);

    const handleLikeClick = () => setIsLiked(!isLiked);

    return (
        <>
            <div className="flex flex-col relative w-full bg-white/95 rounded-[30px]">
                <div className="flex flex-row justify-between items-center mr-5">
                    <div className="flex flex-row items-center ml-2 mt-2 gap-5">
                        <div className="w-12 h-12 bg-[#af5f5f] rounded-[100px]"></div>
                        <div className="flex flex-col">
                            <Link
                                href={`/user/${post.userId}`}
                                className="h-[21px] text-black text-xl font-semibold font-['Inter']"
                            >
                                {post.username}
                            </Link>
                            <span className="h-[29px] text-black/50 text-base font-extralight font-['Inter']">
                                {formatDate(post.timestamp)}
                            </span>
                        </div>
                    </div>
                    <BookmarkIcon />
                </div>
                {post.images.length > 0 && (
                    <div className="h-fitline-clamp-5 overflow-hidden text-black text-base font-normal font-['Inter'] leading-[22px] p-3">
                        {post.images.map((src, idx) => (
                            <a href={src} key={idx} target="_blank">
                                <Image
                                    src={src}
                                    width={100}
                                    height={100}
                                    alt=""
                                />
                            </a>
                        ))}
                    </div>
                )}

                <Link
                    href={`/post/${post.id}`}
                    className="h-[110px] line-clamp-5 overflow-hidden text-black text-base font-normal font-['Inter'] leading-[22px] m-5 mr-7 mb-10"
                >
                    {post.content}
                </Link>
                <div className="flex flex-row gap-5 ml-5">
                    <svg
                        className="cursor-pointer"
                        onClick={handleLikeClick}
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill={isLiked ? "red" : "none"}
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <g id="heart">
                            <path
                                id="Icon"
                                d="M19.5355 5.46436C21.4881 7.41698 21.4881 10.5828 19.5355 12.5354L12.7071 19.3639C12.3166 19.7544 11.6834 19.7544 11.2929 19.3639L4.46447 12.5354C2.51184 10.5828 2.51184 7.41698 4.46447 5.46436C6.0168 3.91202 7.89056 3.43671 9.78125 4.35927C10.5317 4.72543 11.5156 5.46436 12 6.42958C12.4844 5.46436 13.4683 4.72543 14.2187 4.35927C16.1094 3.43671 17.9832 3.91202 19.5355 5.46436Z"
                                stroke="black"
                                strokeWidth="2"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            />
                        </g>
                    </svg>

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
