import React from 'react'
import { UserOnlineIcon } from "./icons/UserOnlineIcon"
import { OnlineUser } from "@/types/user";
import Image from "next/image";
import Link from "next/link";
import useIsMobile from "@/hooks/useIsMobile";

interface UserListProps {
    user: OnlineUser;
    onUserSelect: (user: OnlineUser) => void;
}


const Users = ({ user, onUserSelect }: UserListProps) => {
    const isMobile = useIsMobile();

    const handleUserClick = (e: React.MouseEvent) => {
        e.preventDefault();
        onUserSelect(user);
    };
    return (
        <>
            {isMobile ? <Link href={`/chats/${user.id}`} className="w-full flex flex-row items-center bg-white rounded-3xl cursor-pointer">
                <div className='flex flex-row items-center w-full relative gap-4 p-1'>
                    <Image src={user.image} width={32} height={32} alt="" className='flex justify-center items-center w-9 h-9 border border-black rounded-full'></Image>
                    <div className="flex flex-col">
                        <span >{user.nickname}</span>
                        {/* {user.lastMessage && ( */}
                        <span className="max-w-[240px] text-gray-400 overflow-hidden whitespace-nowrap text-ellipsis inline-block xl:max-w-[110px]">last message message message message message message</span>
                        {/* )} */}
                    </div>
                </div>
                <div className="mr-3"><UserOnlineIcon online={user.online} /></div>
            </Link>
                :
                <div onClick={handleUserClick} className="w-full flex flex-row items-center bg-white rounded-3xl cursor-pointer">
                    <div className='flex flex-row items-center w-full relative gap-4 p-1'>
                        <Image src={user.image} width={32} height={32} alt="" className='flex justify-center items-center w-9 h-9 border border-black rounded-full'></Image>
                        <div className="flex flex-col">
                            <span >{user.nickname}</span>
                            {/* {user.lastMessage && ( */}
                            <span className="max-w-[240px] text-gray-400 overflow-hidden whitespace-nowrap text-ellipsis inline-block xl:max-w-[110px]">last message message message message message message</span>
                            {/* )} */}
                        </div>
                    </div>
                    <div className="mr-3"><UserOnlineIcon online={user.online} /></div>
                </div>
            }
        </>
    )
}

export default Users