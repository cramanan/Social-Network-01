import React from 'react'
import { UserOnlineIcon } from "./icons/UserOnlineIcon"
import { OnlineUser } from "@/types/user";
import Image from "next/image";
import Link from "next/link";

interface UserListProps {
    user: OnlineUser;
}

const Users = ({ user }: UserListProps) => {
    return (
        <>
            <Link key={user.id} href={`/chats/${user.id}`} className="w-full flex flex-row items-center bg-white rounded-3xl">
                <div className='flex flex-row items-center w-full relative gap-4 p-1'>
                    <Image src={user.image} width={32} height={32} alt="" className='flex justify-center items-center w-9 h-9 border border-black rounded-full'></Image>
                    <span >{user.nickname}</span>
                </div>
                <div className="mr-3"><UserOnlineIcon online={user.online} /></div>
            </Link>
        </>
    )
}

export default Users