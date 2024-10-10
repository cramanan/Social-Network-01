import React from 'react'

const ProfileStats = () => {
    return (
        <div className='w-[400px] h-16 bg-white/30 rounded-2xl flex flex-row items-center justify-between px-3'>
            <div className='flex flex-col items-center w-[86px]'>
                <div className='font-bold'>5</div>
                <div className='text-black/50'>Publications</div>
            </div>
            <div className='flex flex-col items-center w-[86px]'>
                <div className='font-bold'>5</div>
                <div className='text-black/50'>Followers</div>
            </div>
            <div className='flex flex-col items-center w-[86px]'>
                <div className='font-bold'>5</div>
                <div className='text-black/50'>Follows</div>
            </div>
            <div className='flex flex-col items-center w-[86px]'>
                <div className='font-bold'>5</div>
                <div className='text-black/50'>Like</div>
            </div>
        </div>
    )
}

export default ProfileStats