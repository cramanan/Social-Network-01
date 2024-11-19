import { User } from "@/types/user"
import Image from "next/image"
import React from 'react'

const FriendInvite = ({ nickname, image }: User) => {
    return (
        <>
            <div className='w-60 flex flex-row justify-between items-center'>
                <Image src={image} alt="" width={40} height={40} ></Image>
                <div className='w-32'>{nickname}</div>
                <div>✓</div>
                <div>X</div>
            </div>
        </>
    )
}

export default FriendInvite