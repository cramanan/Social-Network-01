"use client";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { Group } from "@/types/group";
import { User } from "@/types/user";
import Image from "next/image";
import React, { useEffect, useState } from "react";

function FollowRequests() {
    const [users, setUsers] = useState<User[]>([]);

    const handleRequest = (id: string, foo: "accept" | "decline") => () => {
        fetch(`/api/users/${id}/${foo}-request`, { method: "POST" });
        setUsers(users.filter((u) => u.id !== id));
    };

    useEffect(() => {
        const fetchRequests = async () => {
            const response = await fetch("/api/inbox/follow-requests");
            const data: User[] = await response.json();
            setUsers(data);
        };

        fetchRequests();
    }, []);

    return (
        <div>
            {users.length > 0 ? (
                users.map(({ id, image, nickname }, idx) => (
                    <div key={idx}>
                        <Image src={image} alt="" width={80} height={80} />
                        <span>{nickname}</span>
                        <button onClick={handleRequest(id, "accept")}>✓</button>
                        <button onClick={handleRequest(id, "decline")}>X</button>
                    </div>
                ))
            )
                :
                (
                    <p className="text-center font-bold">
                        No Follow Request(s) found.
                    </p>
                )
            }

        </div>
    );
}

function GroupInvites() {
    const [groups, setGroups] = useState<Group[]>([]);

    const handleRequest = (id: string, foo: "accept" | "decline") => () => {
        fetch(`/api/groups/${id}/${foo}-invite`, { method: "POST" });
        setGroups(groups.filter((group) => group.id !== id));
    };

    useEffect(() => {
        const fetchRequests = async () => {
            const response = await fetch("/api/inbox/group-invites");
            const data: Group[] = await response.json();
            setGroups(data);
        };

        fetchRequests();
    }, []);

    return (
        <div>
            {groups.length > 0 ? (
                groups.map(({ id, name, image }, idx) => (
                    <div key={idx}>
                        <Image src={image} alt="" width={80} height={80} />
                        <span>{name}</span>
                        <button onClick={handleRequest(id, "accept")}>✓</button>
                        <button onClick={handleRequest(id, "decline")}>X</button>
                    </div>
                ))
            )
                :
                (
                    <p className="text-center font-bold">
                        No Group Invite(s) found.
                    </p>
                )
            }

        </div>
    );
}

function GroupRequests() {
    type groupRequest = {
        groupId: Group["id"];
        groupName: Group["name"];
        groupImage: Group["image"];

        userId: User["id"];
        userName: User["nickname"];
        userImage: User["image"];
    };
    const [requests, setRequests] = useState<groupRequest[]>([]);

    // const handleRequest = (id: string, foo: "accept" | "decline") => () => {
    //     fetch(`/api/users/${id}/${foo}-request`, { method: "POST" });
    //     setUsers(users.filter((u) => u.id !== id));
    // };

    useEffect(() => {
        const fetchRequests = async () => {
            const response = await fetch("/api/inbox/group-requests");
            const data: groupRequest[] = await response.json();
            setRequests(data);
        };

        fetchRequests();
    }, []);

    return (
        <div>
            {requests.length > 0 ? (
                requests.map((request, idx) => (
                    <div key={idx} className="w-fit flex flex-col items-center">
                        <div className="flex items-center gap-3">
                            <Image
                                src={request.userImage}
                                alt=""
                                width={80}
                                height={80}
                            />
                            {"=>"}
                            <Image
                                src={request.groupImage}
                                alt=""
                                width={80}
                                height={80}
                            />
                        </div>
                        <span>
                            <a href={`/user/${request.userId}`} target="_blank">
                                {request.userName}
                            </a>{" "}
                            wants to join{" "}
                            <a href={`/group/${request.groupId}`} target="_blank">
                                {request.groupName}
                            </a>
                        </span>
                        {/* <button onClick={handleRequest(id, "accept")}>✓</button>
                        <button onClick={handleRequest(id, "decline")}>
                            X
                        </button> */}
                    </div>
                ))
            )
                :
                (
                    <p className="text-center font-bold">
                        No Group Request(s) found.
                    </p>
                )
            }
        </div>
    );
}

export default function Inbox() {
    const [windows, setWindows] = useState([true, false, false]);
    const changeWindow = (i: number) => () =>
        setWindows((prev) => prev.map(({ }, idx) => idx === i));
    const headers = ["Follow Request", "Group Invite", "Group Request"];
    const content = [FollowRequests, GroupInvites, GroupRequests];

    return (
        <>
            <HomeProfileLayout>
                <nav className="flex gap-3">
                    {headers.map((name, idx) => (
                        <button onClick={changeWindow(idx)} key={idx}>
                            {name}
                        </button>
                    ))}
                </nav>
                {content.map(
                    (Component, idx) => windows[idx] && <Component key={idx} />
                )}
            </HomeProfileLayout>
        </>
    );
}
