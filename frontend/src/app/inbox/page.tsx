"use client";

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
            {users.map(({ id, image, nickname }, idx) => (
                <div key={idx}>
                    <Image src={image} alt="" width={80} height={80} />
                    <span>{nickname}</span>
                    <button onClick={handleRequest(id, "accept")}>✓</button>
                    <button onClick={handleRequest(id, "decline")}>X</button>
                </div>
            ))}
        </div>
    );
}
function GroupInvites() {
    // const [users, setUsers] = useState<User[]>([]);

    // const handleRequest = (id: string, foo: "accept" | "decline") => () => {
    //     fetch(`/api/users/${id}/${foo}-request`, { method: "POST" });
    //     setUsers(users.filter((u) => u.id !== id));
    // };

    useEffect(() => {
        const fetchRequests = async () => {
            const response = await fetch("/api/inbox/group-invites");
            // const data: User[] = await response.json();
            // setUsers(data);
        };

        fetchRequests();
    }, []);

    return (
        <div>
            {/* {users.map(({ id, image, nickname }, idx) => (
                <div key={idx}>
                    <Image src={image} alt="" width={80} height={80} />
                    <span>{nickname}</span>
                    <button onClick={handleRequest(id, "accept")}>✓</button>
                    <button onClick={handleRequest(id, "decline")}>X</button>
                </div>
            ))} */}
        </div>
    );
}

function GroupRequests() {
    // const [users, setUsers] = useState<User[]>([]);

    // const handleRequest = (id: string, foo: "accept" | "decline") => () => {
    //     fetch(`/api/users/${id}/${foo}-request`, { method: "POST" });
    //     setUsers(users.filter((u) => u.id !== id));
    // };

    useEffect(() => {
        const fetchRequests = async () => {
            const response = await fetch("/api/inbox/group-requests");
            // const data: User[] = await response.json();
            // setUsers(data);
        };

        fetchRequests();
    }, []);

    return (
        <div>
            {/* {users.map(({ id, image, nickname }, idx) => (
                <div key={idx}>
                    <Image src={image} alt="" width={80} height={80} />
                    <span>{nickname}</span>
                    <button onClick={handleRequest(id, "accept")}>✓</button>
                    <button onClick={handleRequest(id, "decline")}>X</button>
                </div>
            ))} */}
        </div>
    );
}

export default function Inbox() {
    const [windows, setWindows] = useState([true, false, false]);
    const changeWindow = (i: number) => () =>
        setWindows((prev) => prev.map(({}, idx) => idx === i));
    const headers = ["Follow Requests", "Group Invites", "Group Requests"];
    const content = [FollowRequests, GroupInvites, GroupRequests];

    return (
        <div>
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
        </div>
    );
}
