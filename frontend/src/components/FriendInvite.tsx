'use client'

import { User } from "@/types/user"
import Image from "next/image"
import React from 'react'

const FriendInvite = ({ id, nickname, image }: User) => {
    const acceptRequest = () =>
        fetch(`/api/user/${id}/accept-request`, { method: "POST" });

    const declineRequest = () =>
        fetch(`/api/user/${id}/decline-request`, { method: "POST" });
    console.log(id);

    return (
        <>
            <div className='w-60 flex flex-row justify-between items-center'>
                <Image src={image} alt="" width={40} height={40} ></Image>
                <div className='w-32'>{nickname}</div>
                <button onClick={acceptRequest}>✓</button>
                <button onClick={declineRequest}>X</button>
            </div>
        </>
    )
}

export default FriendInvite