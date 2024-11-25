import React from "react";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import Link from "next/link";
import { Post } from "@/types/post";
import Image from "next/image";
import formatDate from "@/utils/formatDate";

const Media = ({ userId, username, images, timestamp, pfp }: Post) => {
    return (
        <>
            <div className="flex flex-col items-center w-[277px] h-[305px] bg-white rounded-[30px]">
                <Link
                    href={`/user/${userId}`}
                    className="w-[226px] inline-flex items-center gap-3 py-1"
                >
                    <Image src={`${pfp ? ("/") : ("/Default_pfp.jpg")}`}
                        alt=""
                        width={48}
                        height={48}
                        className="flex justify-center items-center w-12 h-12 border border-black rounded-full" />
                    <div className="text-black text-xl font-semibold font-['Inter']">
                        {username}
                    </div>
                </Link>
                {/* {images.map((image, idx) => (
                    <> */}
                <div className="flex justify-center items-center min-w-[226px] min-h-[206px]">
                    <Image
                        src={images[0]}
                        width={250}
                        height={250}
                        alt=""
                        className="object-contain max-w-[226px] max-h-[206px]"
                    />
                </div>
                {/* </> */}
                {/* ))} */}
                {/* <Image src={images[1]} width={236} height={206} alt="" className=' relative bg-[#373333]'></Image> */}
                <div className="w-[226px] flex justify-between gap-5 py-1">
                    <div className="text-black/50 text-base font-extralight font-['Inter']">
                        {formatDate(timestamp)}
                    </div>
                    <LikeIcon />
                    <CommentIcon />
                </div>
            </div>
        </>
    );
};

export default Media;
