"use client";

import React, { useState } from "react";
import { ProfileCircle } from "./icons/ProfileCircle";
import { SendPostIcon } from "./icons/sendPostIcon";
import { CloseIcon } from "./icons/CloseIcon";
import { ImageIcon } from "./icons/ImageIcon";
import { Post } from "@/types/post";
import Image from "next/image";

type PostFields = Pick<Post, "content" | "images">;

export const NewPost = () => {
    const [fields, setFields] = useState<PostFields>({
        content: "",
        images: [],
    });
    const [isModalOpen, setIsModalOpen] = useState(false);

    const toggleModal = () => setIsModalOpen(!isModalOpen);
    const handleTextChange = (e: React.ChangeEvent<HTMLTextAreaElement>) =>
        setFields({ ...fields, content: e.target.value });

    const handleImagesChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const files = e.target.files;
        if (!files) return;

        const images: string[] = [];
        for (const file of files) images.push(URL.createObjectURL(file));

        setFields({ ...fields, images });
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const body = new FormData(e.currentTarget);
        body.append("content", fields.content);
        const response = await fetch("/api/post", { method: "POST", body });

        if (response.ok) toggleModal();
    };

    return (
        <>
            <div className={isModalOpen ? "modal-open" : ""}>
                <div>
                    <button
                        className="px-4 py-2 relative bg-slate-50  hover:bg-slate-300 p-4 rounded-lg shadow-lg flex flex-row items-center gap-4"
                        onClick={toggleModal}
                    >
                        <ProfileCircle />
                        <div className="text-slate-500  font-extralight font-['Inter']">
                            Create your Post
                        </div>
                    </button>
                </div>
            </div>
            {isModalOpen && (
                <form
                    encType="multipart/form-data"
                    onSubmit={handleSubmit}
                    className="fixed top-0 inset-0 flex items-center justify-center backdrop-blur-sm z-50"
                >
                    <div className="border border-white bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7] p-6 rounded-lg shadow-lg  w-1/2 ">
                        <div className="flex justify-between">
                            <h2 className="text-xl text-white font-semibold flex justify-center items-center gap-4 ">
                                <ProfileCircle />
                                New Post
                            </h2>
                            <button
                                className="w-10 h-10 flex justify-center items-center hover:bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7]"
                                onClick={toggleModal}
                            >
                                <CloseIcon />
                            </button>
                        </div>

                        <textarea
                            id="content"
                            className="shadow-lg w-full px-12 py-4 mt-7 rounded-xl  bg-white text-black text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-gray-500 resize-none"
                            value={fields.content}
                            onChange={handleTextChange}
                            placeholder="What's on your mind?"
                        />

                        <div className="flex p-1 gap-2">
                            {fields.images.map((image, idx) => (
                                <Image
                                    key={idx}
                                    src={image}
                                    alt=""
                                    height={100}
                                    width={100}
                                    className="object-contain"
                                />
                            ))}
                        </div>
                        <div className="flex justify-between">
                            <div className="relative flex flex-row gap-2 shadow-xl bg-slate-300  hover:bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7] text-black rounded justify-center items-center px-2 py-1">
                                <input
                                    className="opacity-0 absolute w-full h-full cursor-pointer"
                                    type="file"
                                    name="images"
                                    id="images"
                                    accept="image/jpeg,image/png,image/gif"
                                    onChange={handleImagesChange}
                                    multiple
                                />
                                upload
                                <ImageIcon />
                            </div>
                            <button
                                className=" flex flex-row gap-2 shadow-xl bg-slate-300  hover:bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7] text-black rounded justify-center items-center px-2 py-1"
                                type="submit"
                            >
                                Send Post
                                <SendPostIcon />
                            </button>
                        </div>
                    </div>
                </form>
            )}
        </>
    );
};
