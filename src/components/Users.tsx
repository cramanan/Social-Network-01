import React from 'react'
import { UserRect } from './icons/UserRect'
import { UserOnline } from './icons/UserOnline'
import { UserProfil } from './icons/UserProfil'

const Users = () => {
    return (
        <div className='flex'>
            <UserRect />
            <div className='flex justify-between items-center w-40 relative z-10 -ml-[170px] -mt-0.5'>
                <div className='flex justify-center items-center w-9 h-9 bg-black rounded-full'>
                    <UserProfil />
                </div>
                <span className='-ml-10'>user</span>
                <UserOnline />
            </div>
        </div>
    )
}

export default Users