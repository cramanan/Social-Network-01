import Image from "next/image";
import React, { useEffect, useState } from "react";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import Comment from "./Comment";
import Link from "next/link";
import formatDate from "@/utils/formatDate";
import { Comment as CommentType, Post } from "@/types/post";
import { NewComment } from "./NewComment";

export const PostMedia = ({ post }: { post: Post }) => {
    const [isLiked, setIsLiked] = useState(false);
    const handleLikeCLick = () => setIsLiked(!isLiked);
    const [allComments, setAllComments] = useState<CommentType[]>([]);

    useEffect(() => {
        const fetchComments = async () => {
            const response = await fetch(`/api/posts/${post.id}/comments`);
            const data: CommentType[] = await response.json();

            setAllComments(data);
        };

        fetchComments();
    }, [post.id]);
    return (
        <>
            <div className="flex flex-col relative w-full bg-white/95 px-5 py-2 md:flex-row xl:rounded-[30px]">
                <div className=" flex flex-col gap-3 pr-2">
                    <div className="flex justify-center min-w-[300px]">
                        <a
                            href={post.images[0]}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="h-fit"
                        >
                            <Image
                                src={post.images[0]}
                                width={500}
                                height={500}
                                alt=""
                                className="max-h-[300px] w-auto h-auto object-contain"
                                priority
                            ></Image>
                        </a>
                    </div>
                    <div className=" flex flex-row gap-10">
                        <div onClick={handleLikeCLick} className="flex gap-2">
                            <LikeIcon isLiked={isLiked} />
                            0
                        </div>
                        <div className="flex gap-2">
                            <CommentIcon />
                            {allComments.length}
                        </div>
                    </div>
                </div>

                <div className="flex flex-col w-full pl-2">
                    <div className="flex flex-row justify-between items-center">
                        <div className="flex flex-row justify-center items-center gap-2">
                            <Link href={`/user/${post.userId}`}>
                                <Image
                                    src={post.userImage}
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
                    </div>

                    <p className="w-full max-h-[150px] my-1 overflow-y-auto whitespace-pre-wrap md:h-[150px]">
                        {post.content}
                    </p>

                    <div className="max-h-[108px] overflow-y-auto bg-black/10 my-2">
                        {allComments.map((comment, idx) => (
                            <Comment key={idx} {...comment} />
                        ))}
                    </div>

                    <NewComment {...post} />
                </div>
            </div>
        </>
    );
};
