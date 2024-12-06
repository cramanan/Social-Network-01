import React from "react";
import { LikeIcon } from "./icons/LikeIcon";
import { Comment as CommentType } from "@/types/post";
import Image from "next/image";
import { useAuth } from "@/hooks/useAuth";
import Link from "next/link";

const Comment = ({ userId, content, username, userImage, image }: CommentType) => {
    const { user } = useAuth()
    return (
        <div
            className="flex items-center justify-between relative w-full h-[54px] bg-[#f6f6f6]/0 px-2"
            aria-label="Comment"
        >
            <div className="flex flex-row items-center">
                <div className="m-2">
                    <Image
                        src={userImage}
                        width={38}
                        height={38}
                        alt=""
                        className="w-38 h-auto bg-[#b53695] rounded-[100px]"
                    />
                </div>
                <div className="flex flex-col">
                    <Link
                        href={`${user?.id === userId ? `/profile` : `/user/${userId}`}`}
                        className="h-[21px] text-black text-[15px] font-semibold font-['Inter']">
                        {username}
                    </Link>
                    { }
                    <span className="flex items-center gap-2 text-black text-[13px] font-normal font-['Inter']">
                        {image &&
                            <a href={image}>
                                <Image
                                    src={image}
                                    width={28}
                                    height={28}
                                    alt=""
                                    className="object-contain w-7 h-7"
                                />
                            </a>
                        }
                        {content}
                    </span>
                </div>
            </div>
            <LikeIcon />
        </div>
    );
};

export default Comment;
