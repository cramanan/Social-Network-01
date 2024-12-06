"use client";

import React, { useEffect, useState } from "react";
import { ProfileCircle } from "./icons/ProfileCircle";
import { SendPostIcon } from "./icons/sendPostIcon";
import { CloseIcon } from "./icons/CloseIcon";
import { ImageIcon } from "./icons/ImageIcon";
import { Post } from "@/types/post";
import Image from "next/image";
import { useAuth } from "@/hooks/useAuth";
import { User } from "@/types/user";

type PostFields = Pick<
    Post,
    "groupId" | "content" | "images" | "privacyLevel" | "selectedUserIds"
>;

const privacyLevels: Post["privacyLevel"][] = [
    "public",
    "almost_private",
    "private",
];

export const NewPost = ({ groupId }: { groupId: string | null }) => {
    const { user } = useAuth();
    const [fields, setFields] = useState<PostFields>({
        content: "",
        images: [],
        groupId,
        privacyLevel: "public",
        selectedUserIds: [],
    });
    const [userIds, setUserIds] = useState<User[]>([]);

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
        body.append("data", JSON.stringify(fields));
        const response = await fetch("/api/posts", { method: "POST", body });

        if (response.ok) toggleModal();
    };

    useEffect(() => {
        if (fields.privacyLevel !== "private") return;
        (async () => {
            const response = await fetch("api/profile/followers");
            const data: User[] = await response.json();

            setUserIds(data);
        })();
    }, [fields.privacyLevel]);

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
                <>
                    <div className="-inset-full fixed bg-black/10 backdrop-blur-sm z-40"></div>
                    <form
                        encType="multipart/form-data"
                        onSubmit={handleSubmit}
                        className="fixed top-0 inset-0 flex items-center justify-center z-50"
                    >
                        <div className="w-4/5 border border-white bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7] p-6 rounded-lg shadow-lg xl:w-1/2 ">
                            <div className="flex justify-between">
                                <h2 className="text-xl text-white font-semibold flex justify-center items-center gap-4 ">
                                    <Image
                                        src={user?.image ?? "/Default_pfp.jpg"}
                                        width={40}
                                        height={40}
                                        alt=""
                                        className="w-10 h-10 rounded-full"
                                    />
                                    New Post
                                </h2>
                                <button
                                    className="w-10 h-10 flex justify-center items-center hover:bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7]"
                                    onClick={toggleModal}
                                >
                                    <CloseIcon />
                                </button>
                            </div>
                            <ul className="flex justify-evenly mt-5">
                                {privacyLevels.map((level) => (
                                    <li key={level} className="flex gap-2">
                                        <input
                                            key={level}
                                            type="radio"
                                            id={level}
                                            name="privacyLevel"
                                            value={level}
                                            checked={
                                                fields.privacyLevel === level
                                            }
                                            onChange={() =>
                                                setFields({
                                                    ...fields,
                                                    privacyLevel: level,
                                                })
                                            }
                                        />

                                        <label htmlFor={level}>{level}</label>
                                    </li>
                                ))}
                            </ul>
                            {fields.privacyLevel === "private" &&
                                (<ul>
                                    {userIds.length > 0 ?
                                        (userIds.map((user, idx) => (
                                            <li key={idx}>
                                                <input
                                                    type="checkbox"
                                                    value={user.id}
                                                    onChange={(e) => {
                                                        const update = e.target.checked
                                                            ? [
                                                                ...fields.selectedUserIds,
                                                                user.id,
                                                            ]
                                                            : fields.selectedUserIds.filter(
                                                                (id) => id != user.id
                                                            );

                                                        setFields({
                                                            ...fields,
                                                            selectedUserIds: update,
                                                        });
                                                    }}
                                                />
                                                <label>{user.nickname}</label>
                                            </li>
                                        )))
                                        :
                                        (
                                            <p className="text-center font-bold">
                                                No follower(s) found.
                                            </p>
                                        )
                                    }
                                </ul>)
                            }


                            <textarea
                                id="content"
                                className="shadow-lg w-full px-12 py-4 mt-5 rounded-xl  bg-white text-black text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-gray-500 resize-none"
                                value={fields.content}
                                onChange={handleTextChange}
                                placeholder="What's on your mind?"
                            />
                            <ul className="flex gap-1 mb-1">
                                {fields.images.map((image, idx) => (
                                    <li key={idx}>
                                        <Image
                                            src={image}
                                            height={100}
                                            width={100}
                                            alt=""
                                        />
                                    </li>
                                ))}
                            </ul>
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
                </>
            )}
        </>
    );
};
