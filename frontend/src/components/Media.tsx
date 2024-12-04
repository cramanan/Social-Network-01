import React from "react";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import Link from "next/link";
import { Post } from "@/types/post";
import Image from "next/image";
import formatDate from "@/utils/formatDate";

const Media = ({ userId, username, images, timestamp, userImage }: Post) => {
    return (
        <>
            <div className="flex flex-col items-center w-[277px] h-[305px] bg-white rounded-[30px]">
                <Link
                    href={`/user/${userId}`}
                    className="w-[226px] inline-flex items-center gap-3 py-1"
                >
                    <Image src={`${userImage}`}
                        alt=""
                        width={48}
                        height={48}
                        className="flex justify-center items-center w-12 h-12 border border-black rounded-full" />
                    <div className="text-black text-xl font-semibold font-['Inter']">
                        {username}
                    </div>
                </Link>
                <div className="flex justify-center items-center min-w-[226px] min-h-[206px]">
                    <Image
                        src={images[0]}
                        width={250}
                        height={250}
                        alt=""
                        className="object-contain max-w-[226px] max-h-[206px]"
                    />
                </div>
                <div className="w-[226px] flex justify-between items-center gap-5 py-1">
                    <div className="text-black/50 text-sm font-extralight font-['Inter']">
                        {formatDate(timestamp)}
                    </div>
                    <div className="flex gap-1">
                        <LikeIcon />
                        0
                    </div>
                    <div className="flex gap-1">
                        <CommentIcon />
                        0 {/* TODO: change to dynamic  */}
                    </div>

                </div>
            </div>
        </>
    );
};

export default Media;
