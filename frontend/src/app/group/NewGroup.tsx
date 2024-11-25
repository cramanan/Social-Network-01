import { Group } from "@/types/group";
import { StrictOmit } from "@/utils/types";
import Image from "next/image";
import React, { ChangeEvent, FormEvent, useState } from "react";

type GroupFields = StrictOmit<Group, "id" | "timestamp">;

export default function NewGroup() {
    const [formState, setFormState] = useState<GroupFields>({
        name: "",
        description: "",
        image: "/group-default.png",
    });

    const handleFileUpload = (e: ChangeEvent<HTMLInputElement>) => {
        const files = e.target.files;
        if (!files) return;

        const image = URL.createObjectURL(files[0]);

        setFormState({ ...formState, image });
    };

    const onSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const body = new FormData(e.currentTarget);
        body.append("data", JSON.stringify(formState));
        fetch("/api/create/group", { method: "POST", body });
    };

    return (
        <form
            onSubmit={onSubmit}
            className="flex flex-col w-full items-center p-2 shadow-xl xl:flex-row xl:gap-2"
        >
            <div className="w-full flex flex-row justify-evenly items-center xl:gap-2 xl:w-fit">
                <label htmlFor="images">
                    <Image
                        src={formState.image}
                        alt=""
                        height={64}
                        width={64}
                        className="max-w-16 max-h-16 rounded-full object-cover"
                        priority
                    />
                </label>
                <input
                    type="file"
                    name="images"
                    id="images"
                    onChange={handleFileUpload}
                    className="absolute hidden"
                />

                <div className="flex flex-col items-center">
                    <label htmlFor="name">Group Name</label>
                    <input
                        type="text"
                        id="name"
                        onChange={(e) =>
                            setFormState({
                                ...formState,
                                name: e.target.value,
                            })
                        }
                    />
                </div>
            </div>

            <div className="flex flex-col w-full">
                <label htmlFor="description">Description</label>
                <input
                    type="text"
                    id="description"
                    onChange={(e) =>
                        setFormState({
                            ...formState,
                            description: e.target.value,
                        })
                    }
                />
            </div>

            <button type="submit">Create Group</button>
        </form>
    );
}
