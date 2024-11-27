import Image from "next/image"
import React, { useState } from 'react'
import { LikeIcon } from "./icons/LikeIcon"
import { CommentIcon } from "./icons/CommentIcon"
import { BookmarkIcon } from "./icons/BookmarkIcon"
import Comment from "./Comment"
import { Post } from "@/types/post"
import Link from "next/link"
import formatDate from "@/utils/formatDate"

export const PostMedia = ({ post }: { post: Post }) => {
    const [isLiked, setIsLiked] = useState(false)

    const handleLikeCLick = () => setIsLiked(!isLiked)
    return (
        <>
            <div className="flex flex-col relative w-full bg-white/95 px-5 py-2 md:flex-row xl:rounded-[30px]">
                <div className=" flex flex-col gap-3 pr-2">
                    <div className="flex justify-center min-w-[300px]">
                        <a href={post.images[0]} target="_blank" rel="noopener noreferrer" className="h-fit">
                            <Image src={post.images[0]} width={500} height={500} alt="" className="max-h-[300px] w-auto h-auto object-contain"></Image>
                        </a>
                    </div>
                    <div className=" flex flex-row gap-10">
                        <div onClick={handleLikeCLick}>
                            <LikeIcon isLiked={isLiked} />
                        </div>
                        <CommentIcon />
                    </div>
                </div>

                <div className="flex flex-col w-full pl-2">
                    <div className="flex flex-row justify-between items-center">
                        <div className="flex flex-row justify-center items-center gap-2">
                            <Link href={`/user/${post.userId}`}>
                                <Image
                                    src={`${post.pfp ? ("/") : ("/Default_pfp.jpg")}`}
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

                    <p className="w-full max-h-[150px] overflow-scroll no-scrollbar my-1 md:h-[150px]">{post.content}</p>

                    <div className="bg-black/10 mt-2 mb-5">
                        <Comment />
                        <Comment />
                    </div>

                    <div className="h-[58px] pl-px pr-3 pt-[11px] pb-[7px] bg-[#f2eeee] rounded-[10px] justify-between items-center inline-flex">
                        <div className='flex flex-row items-center gap-2'>
                            <div className="w-[44px] h-[40px] relative">emote</div>
                            <input type='text' placeholder='Enter your comment' className="w-[300px] h-[30px] text-black text-xl font-extralight font-['Inter'] bg-white/0"></input>
                        </div>
                        <div className="self-stretch pl-[11px] pr-3 pt-[5px] bg-gradient-to-t from-[#e1d3eb] via-[#6f46c0] to-[#e0d3ea] rounded-[30px] justify-center items-center inline-flex">
                            <button className="h-[25px] text-center text-black text-[15px] font-medium font-['Inter']">Send</button>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}
