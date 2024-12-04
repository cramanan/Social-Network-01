import React, { ChangeEvent, FormEvent, useState } from 'react'
import { Comment as CommentType, Post } from "@/types/post";
import { ImageIcon } from "./icons/ImageIcon";
import Image from "next/image";

type CommentFields = Pick<CommentType, "content" | "image">;

const defaultComment = {
    content: "",
    image: "",
};

export const NewComment = ({ id }: Post) => {
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

    const submitComment = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formdata = new FormData(e.currentTarget);
        formdata.append(
            "data",
            JSON.stringify({ content: newComment.content })
        );
        console.log(formdata);
        try {
            const response = await fetch(`/api/posts/${id}/comments/`, {
                method: "POST",
                body: formdata,
            });
            if (response.ok) setComment(defaultComment);
        } catch (error) {
            console.error(error);
        }
    };
    return (
        <>
            <form
                onSubmit={submitComment}
                className="px-3 pt-[11px] pb-[7px] bg-[#f2eeee] rounded-[10px] gap-2 items-center inline-flex mx-5 my-2"
            >
                <div className="w-full flex flex-row items-center gap-2">
                    <label
                        htmlFor="images"
                        className="w-fit text-center cursor-pointer"
                    >
                        <ImageIcon />
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
                            className="w-full px-2"
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
        </>
    )
}
