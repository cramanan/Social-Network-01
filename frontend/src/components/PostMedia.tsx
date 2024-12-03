import Image from "next/image";
import React, { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { LikeIcon } from "./icons/LikeIcon";
import { CommentIcon } from "./icons/CommentIcon";
import Comment from "./Comment";
import Link from "next/link";
import formatDate from "@/utils/formatDate";
import { Comment as CommentType, Post } from "@/types/post";

type CommentFields = Pick<CommentType, "content" | "image">;

const defaultComment = {
    content: "",
    image: "",
};

export const PostMedia = ({ post }: { post: Post }) => {
    const [isLiked, setIsLiked] = useState(false);

    const handleLikeCLick = () => setIsLiked(!isLiked);

    const [newComment, setComment] = useState<CommentFields>(defaultComment);

    const changeCommentContent = (e: ChangeEvent<HTMLInputElement>) =>
        setComment({ ...newComment, content: e.target.value });
    const changeCommentImages = (e: ChangeEvent<HTMLInputElement>) => {
        if (!e.target.files) return;
        setComment({
            ...newComment,
            image: URL.createObjectURL(e.target.files[0]),
        });
    };

    const [allComments, setAllComments] = useState<CommentType[]>([]);

    const submitComment = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formdata = new FormData(e.currentTarget);
        formdata.append(
            "data",
            JSON.stringify({ content: newComment.content })
        );
        console.log(formdata);
        try {
            const response = await fetch(`/api/posts/${post.id}/comments/`, {
                method: "POST",
                body: formdata,
            });
            if (response.ok) setComment(defaultComment);
        } catch (error) {
            console.error(error);
        }
    };

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
                            ></Image>
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

                    <p className="w-full max-h-[150px] overflow-scroll no-scrollbar my-1 md:h-[150px]">
                        {post.content}
                    </p>

                    <div className="h-[108px] overflow-scroll no-scrollbar bg-black/10 mt-2 mb-5">
                        {allComments.map((comment, idx) => (
                            <Comment key={idx} {...comment} />
                        ))}
                    </div>

                    <form
                        onSubmit={submitComment}
                        className="px-3 pt-[11px] pb-[7px] bg-[#f2eeee] rounded-[10px] gap-2 items-center inline-flex mx-5 my-2"
                    >
                        <div className="w-full flex flex-row items-center gap-2">
                            <label
                                htmlFor="images"
                                className="w-fit text-center cursor-pointer"
                            >
                                Send image
                            </label>
                            <input
                                name="images"
                                id="images"
                                type="file"
                                className="hidden"
                                accept="image/jpeg,image/png,image/gif"
                                onChange={changeCommentImages}
                            />
                            <div className="w-full h-[30px] text-black text-xl font-extralight font-['Inter'] bg-white/0">
                                <div>
                                    {newComment.image && (
                                        <Image
                                            src={newComment.image}
                                            alt=""
                                            width={40}
                                            height={40}
                                        />
                                    )}
                                </div>
                                <input
                                    value={newComment.content}
                                    type="text"
                                    placeholder="Enter your newComment"
                                    className="w-full"
                                    onChange={changeCommentContent}
                                />
                            </div>
                        </div>
                        <div className="self-stretch pl-[11px] pr-3 pt-[5px] bg-gradient-to-t from-[#e1d3eb] via-[#6f46c0] to-[#e0d3ea] rounded-[30px] justify-center items-center inline-flex">
                            <button
                                type="submit"
                                className="h-[25px] text-center text-black text-[15px] font-medium font-['Inter']"
                            >
                                Send
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </>
    );
};
