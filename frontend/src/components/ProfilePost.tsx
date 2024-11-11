import React from "react";
import Comment from "./Comment";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import { Post } from "@/types/post";

const ProfilePost = ({ content, timestamp }: Post) => {
    return (
        <>
            <div className="flex flex-row">
                <div className="w-[800px] h-[300px] bg-white rounded-l-[30px] flex flex-col justify-between">
                    <div className="resize-none w-full h-44 py-4 px-7 rounded-tl-[30px]">
                        {content}
                    </div>
                    <div className="flex flex-row justify-between p-7 pb-5">
                        <div className="text-black/50">{timestamp}</div>
                        <div className="flex flex-row gap-20">
                            <LikeIcon />
                            <CommentIcon />
                        </div>
                    </div>
                </div>
                <div className="w-[300px] h-[300px] bg-white/20 rounded-r-[30px] flex flex-col justify-center pl-2 pr-10">
                    <div className="h-[270px] overflow-scroll no-scrollbar">
                        <Comment />
                        <Comment />
                    </div>
                </div>
            </div>
        </>
    );
};

export default ProfilePost;
