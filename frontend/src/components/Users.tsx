import React from 'react'
import { UserProfil } from './icons/UserProfil'
import { UserOnlineIcon } from "./icons/UserOnlineIcon"

const Users = () => {
    return (
        <>
            <div className="w-full flex flex-row items-center bg-white rounded-3xl">
                <div className='flex flex-row items-center w-full relative gap-4 p-1'>
                    <div className='flex justify-center items-center w-9 h-9 bg-black rounded-full'>
                        <UserProfil />
                    </div>
                    <span >user</span>
                </div>
                <div className="mr-3"><UserOnlineIcon /></div>
            </div>
        </>
    )
}

export default Users