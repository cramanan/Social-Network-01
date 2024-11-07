"use client";

import { EditableUser } from "@/types/user";
import Image from "next/image";
import React, { ChangeEvent, FormEvent, useState } from "react";

export default function UserInfos({
    nickname,
    firstName,
    lastName,
    image,
    aboutMe,
    isPrivate,
}: EditableUser) {
    const [formState, setFormState] = useState<EditableUser>({
        nickname,
        firstName,
        lastName,
        image,
        aboutMe,
        isPrivate,
    });

    const onChange = (e: ChangeEvent<HTMLInputElement>) => {
        const files = e.target.files;
        if (!files) return;

        const image = URL.createObjectURL(files[0]);
        setFormState({ ...formState, image });
    };

    const onSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const body = new FormData(e.currentTarget);
        fetch("/api/auth", { method: "PATCH", body });
    };

    return (
        <form className="w-fit flex flex-col" onSubmit={onSubmit}>
            <div className="w-fit relative flex items-center">
                <label
                    htmlFor="image"
                    className="h-full w-full flex items-center justify-center absolute"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        width={24}
                        height={24}
                        fill="currentColor"
                    >
                        <path d="M4 19H20V12H22V20C22 20.5523 21.5523 21 21 21H3C2.44772 21 2 20.5523 2 20V12H4V19ZM13 9V16H11V9H6L12 3L18 9H13Z"></path>
                    </svg>
                </label>
                <input
                    type="file"
                    name="image"
                    id="image"
                    onChange={onChange}
                    className="absolute hidden"
                />
                <Image
                    src={formState.image}
                    alt=""
                    width={100}
                    height={100}
                    className="w-24 h-24 rounded-full object-cover"
                />
            </div>
            <label htmlFor="nickname">Nickname</label>
            <input
                type="text"
                id="nickname"
                defaultValue={nickname}
                onChange={(e) =>
                    setFormState({ ...formState, nickname: e.target.value })
                }
            />
            <label htmlFor="first-name">FirstName</label>
            <input
                type="text"
                id="first-name"
                defaultValue={firstName}
                onChange={(e) =>
                    setFormState({ ...formState, firstName: e.target.value })
                }
            />
            <label htmlFor="last-name">Last Name</label>
            <input
                type="text"
                id="last-name"
                defaultValue={lastName}
                onChange={(e) =>
                    setFormState({ ...formState, lastName: e.target.value })
                }
            />
            <label htmlFor="about-me">About Me</label>
            <input
                type="text"
                id="about-me"
                defaultValue={aboutMe ?? ""}
                placeholder="Write something about yourself"
                onChange={(e) =>
                    setFormState({ ...formState, aboutMe: e.target.value })
                }
            />
            <label htmlFor="is-private">Set account as private</label>
            <input
                type="checkbox"
                id="is-private"
                defaultValue={`${isPrivate}`}
                onChange={(e) =>
                    setFormState({ ...formState, isPrivate: e.target.checked })
                }
            />
            <button type="submit">Save changes</button>
        </form>
    );
}
