import React, { useEffect, useState } from "react";
import { User } from "@/types/user";
import Image from "next/image";

const FollowInviteList = () => {
    const [users, setUsers] = useState<User[]>([]);

    const handleRequest = (id: string, foo: "accept" | "decline") => () => {
        fetch(`/api/user/${id}/${foo}-request`, { method: "POST" });
        setUsers(users.filter((u) => u.id !== id));
    };

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("/api/follow-requests");
            const data: User[] = await response.json();

            setUsers(data);
        };
        fetchUsers();
    }, []);

    return (
        <div
            id="followInviteList"
            className="relative flex flex-col w-full h-full xl:w-fit xl:h-fit xl:bg-white/25 xl:rounded-[30px] xl:px-2 xl:py-5"
        >
            <h2 className="text-4xl text-white text-center py-5 xl:sr-only">
                Follow Request List
            </h2>

            <ul className="flex flex-col h-fit items-center gap-3 mx-2 overflow-scroll no-scrollbar xl:max-h-[68vh] xl:gap-1">
                {users.length > 0 ? (
                    users.map(({ id, image, nickname }, idx) => (
                        <li
                            key={idx}
                            className="w-60 flex flex-row justify-between items-center "
                        >
                            <Image
                                src={image}
                                alt=""
                                width={40}
                                height={40}
                            ></Image>
                            <div className="w-32">{nickname}</div>
                            <button onClick={handleRequest(id, "accept")}>
                                ✓
                            </button>
                            <button onClick={handleRequest(id, "decline")}>
                                X
                            </button>
                        </li>
                    ))
                ) : (
                    <p className="text-center font-bold">No invite(s) found.</p>
                )}
            </ul>
        </div>
    );
};

export default FollowInviteList;
