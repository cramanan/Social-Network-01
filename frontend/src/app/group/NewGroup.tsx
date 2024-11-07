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
        <form onSubmit={onSubmit} className="flex flex-col items-start">
            <label htmlFor="images">
                <Image
                    src={formState.image}
                    alt=""
                    height={96}
                    width={96}
                    className="h-24 w-24 rounded-full object-cover"
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
            <button type="submit">Create Group</button>
        </form>
    );
}
