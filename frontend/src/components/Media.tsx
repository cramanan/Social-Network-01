import React from "react";
import { ProfileCircle } from "./icons/ProfileCircle";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import Link from "next/link";
import { Post } from "@/types/post";
import Image from "next/image";

const Media = ({ username, images, timestamp }: Post) => {
    return (
        <>
            <div className="flex flex-col items-center w-[277px] h-[305px] bg-white rounded-[30px]">
                <Link
                    href={`/profile`}
                    className="w-[226px] inline-flex items-center gap-3 py-1"
                >
                    <ProfileCircle />
                    <div className="text-black text-xl font-extralight font-['Inter']">
                        {username}
                    </div>
                </Link>
                {images.map((image, idx) => (
                    <>
                        <div className="flex justify-center items-center min-w-[226px] min-h-[206px]">
                            <Image
                                key={idx}
                                src={image}
                                width={250}
                                height={250}
                                alt=""
                                className="object-contain max-w-[226px] max-h-[206px]"
                            />
                        </div>
                    </>
                ))}
                {/* <Image src={images[1]} width={236} height={206} alt="" className=' relative bg-[#373333]'></Image> */}
                <div className="w-[226px] flex justify-between gap-5 py-1">
                    <div className="text-black/50 text-[11px] font-extralight font-['Inter']">
                        {timestamp}
                    </div>
                    <LikeIcon />
                    <CommentIcon />
                </div>
            </div>
        </>
    );
};

export default Media;
