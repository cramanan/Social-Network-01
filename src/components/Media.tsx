import React from 'react'
import { Like } from './icons/Like'
import { NewComment } from './icons/NewComment'
import { ProfileCircle } from './icons/ProfileCircle'

const Media = () => {
  return (
    <div className='flex flex-col items-center w-[307px] h-[335px] bg-white rounded-[30px]'>
      <div className='w-[256px] inline-flex items-center gap-3 mb-1 mt-2'>
        <ProfileCircle />
        <div className="text-black text-xl font-extralight font-['Inter']">User</div>
      </div>
      <div className='w-[256px] h-[233px] relative bg-[#373333]'></div>
      <div className='w-[256px] flex justify-between gap-5 mt-1'>
        <div className="text-black/50 text-[11px] font-extralight font-['Inter']">Friday 6 september 16:03</div>
        <Like />
        <NewComment />
      </div>
    </div>
  )
}

export default Media