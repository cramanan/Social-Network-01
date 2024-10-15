import React from "react";
import { Bookmark } from "./icons/Bookmark";
import { NewComment } from "./icons/NewComment";
import Comment from "./Comment";
import { Like } from "./icons/Like";
import { Post } from "@/types/post";
import Link from "next/link";

const PostComponent = ({ post }: { post: Post }) => {
    const onLike = () => fetch(`/api/post/${post.id}/like`, { method: "POST" });

    return (
        <div className="flex flex-col relative w-full bg-white/95 rounded-[30px]">
            <div className="flex flex-row justify-between items-center mr-5">
                <div className="flex flex-row items-center ml-2 mt-2 gap-5">
                    <div className="w-12 h-12 bg-[#af5f5f] rounded-[100px]"></div>
                    <div className="flex flex-col">
                        <Link
                            href={`/user/${post.userId}`}
                            className="h-[21px] text-black text-xl font-semibold font-['Inter']"
                        >
                            {post.userId}
                        </Link>
                        <span className="h-[29px] text-black/50 text-base font-extralight font-['Inter']">
                            {new Date(post.timestamp).toLocaleDateString(
                                "en-US"
                            )}
                        </span>
                    </div>
                </div>
                <Bookmark />
            </div>
            <div className="h-[110px] line-clamp-5 overflow-hidden text-black text-base font-normal font-['Inter'] leading-[22px] m-5 mr-7 mb-10">
                {post.content}
            </div>
            <div className="flex flex-row gap-5 ml-5">
                <span className="cursor-pointer" onClick={onLike}>
                    <Like />
                </span>
                <NewComment />
            </div>
            <div className="mb-5 mt-1 ml-5 mr-10">
                <Comment />
                <Comment />
            </div>
            <div className="text-center text-black text-sm font-medium font-['Inter'] mb-2 cursor-pointer">
                See more
            </div>
        </div>
    );
};

export default PostComponent;
