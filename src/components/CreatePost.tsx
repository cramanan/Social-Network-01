import React from 'react'
import { ProfileCircle } from './icons/ProfileCircle'

const CreatePost = () => {
    return (
        <div className='flex flex-row items-center h-[70px] relative bg-white/95 rounded-[30px] w-[500px] justify-between'>
            <div>
                <div className='flex flex-row gap-5 ml-5'>
                    <ProfileCircle />
                    <input placeholder='Create your post' className='md:w-[250px] bg-white/0 outline-none'></input>
                </div>
            </div>
            <button className='flex items-center justify-center mr-5 w-[85px] bg-gradient-to-tr from-[#4821f9] via-[#6f46c0] to-[#e0d3ea] rounded-[30px]'>
                <span className=" text-black text-base font-semibold font-['Inter']">post</span>
            </button>
        </div>
    )
}

export default CreatePost