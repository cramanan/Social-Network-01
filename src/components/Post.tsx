"use client";

import React, { useState } from "react";
import Comment from "./Comment";
import { RedLikeIcon } from "./icons/RedLikeIcon";
import { BookmarkIcon } from "./icons/BookmarkIcon";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";

const Post = () => {
    const [isLiked, setIsLiked] = useState(false);

    const handleLikeClick = () => {
        setIsLiked(!isLiked);
    };

    return (
        <>
            <div className="flex flex-col relative w-full bg-white/95 rounded-[30px]">
                <div className="flex flex-row justify-between items-center mr-5">
                    <div className="flex flex-row items-center ml-2 mt-2 gap-5">
                        <div className="w-12 h-12 bg-[#af5f5f] rounded-[100px]"></div>
                        <div className="flex flex-col">
                            <span className="h-[21px] text-black text-xl font-semibold font-['Inter']">
                                User
                            </span>
                            <span className="h-[29px] text-black/50 text-base font-extralight font-['Inter']">
                                Friday 6 september 16:03
                            </span>
                        </div>
                    </div>
                    <BookmarkIcon />
                </div>
                <div className="h-[110px] line-clamp-5 overflow-hidden text-black text-base font-normal font-['Inter'] leading-[22px] m-5 mr-7 mb-10">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                    Nulla lectus enim, dignissim id consectetur ut, congue sit
                    amet libero. Nullam in lorem mollis, sollicitudin est et,
                    ornare augue.
                    <br />
                    Suspendisse risus est, porttitor vitae orci eget, sagittis
                    interdum est. Nunc turpis nisl, vestibulum non condimentum
                    eget, eleifend gravida justo. Nunc turpis nisl, vestibulum
                    non condimentum eget, eleifend gravida justo.
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

export default Post;
