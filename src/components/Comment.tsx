import React from 'react'
import { Like } from './icons/Like'

const Comment = () => {
    return (
        <div className="flex items-center justify-between relative w-full h-[54px] bg-[#f6f6f6]/0">
            <div className='flex flex-row items-center'>
                <div className='m-2'>
                    <div className="w-[41px] h-10 bg-[#b53695] rounded-[100px]"></div>
                </div>
                <div className='flex flex-col'>
                    <span className="h-[21px] text-black text-[15px] font-semibold font-['Inter']">Name</span>
                    <span className=" text-black text-[13px] font-normal font-['Inter']">sss</span>
                </div>
            </div>
            <div className='mr-2'><Like /></div>
        </div>
    )
}

export default Comment