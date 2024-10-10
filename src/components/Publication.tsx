import React from 'react'
import { ProfileCircle } from './icons/ProfileCircle'
import { Like } from './icons/Like'
import { NewComment } from './icons/NewComment'
import Comment from './Comment'

const Publication = () => {
    return (
        <div className='w-[900px] h-[500px] bg-[#e8e8e8]/95 rounded-[30px] flex'>
            <div className='w-[350px] h-[450    ] relative bg-[#373333] m-4'></div>
            <div className='m-4 ml-0 flex flex-col justify-between'>
                <div className='flex gap-3 items-center'>
                    <div><ProfileCircle /></div>
                    <div className=" text-black text-xl font-extralight font-['Inter']">user</div>
                </div>
                <div className='w-[500px] h-[120px] line-clamp-5 text-ellipsis overflow-hidden bg-white rounded-xl mt-2 pl-2'>
                    lorem ipsum Exsistit autem hoc loco quaedam quaestio subdifficilis, num quando amici novi, digni amicitia, veteribus sint anteponendi, ut equis vetulis teneros anteponere solemus. Indigna homine dubitatio! lorem ipsum Exsistit autem hoc loco quaedam quaestio subdifficilis, num quando amici novi, digni amicitia, veteribus sint anteponendi, ut equis vetulis teneros anteponere solemus. Indigna homine dubitatio!
                </div>
                <div className='flex items-center justify-between ml-1 mt-3 mr-2'>
                    <div className="text-black/50 text-[11px] font-extralight font-['Inter']">Friday 6 september 16:03</div>
                    <div className='flex gap-7'>
                        <div><Like /></div>
                        <div><NewComment /></div>
                    </div>
                </div>
                <div className='mt-2 h-[162px] overflow-scroll no-scrollbar bg-black/5'>
                    <Comment />
                    <Comment />
                    <Comment />
                </div>
                <div className="h-[58px] pl-px pr-3 pt-[11px] pb-[7px] bg-[#f2eeee] rounded-[10px] justify-between items-center inline-flex">
                    <div className='flex flex-row items-center gap-2'>
                        <div className="w-[44px] h-[40px] relative">emote</div>
                        <input type='text' placeholder='Enter your comment' className="w-[300px] h-[30px] text-black text-xl font-extralight font-['Inter'] bg-white/0"></input>
                    </div>
                    <div className="self-stretch pl-[11px] pr-3 pt-[5px] bg-gradient-to-t from-[#e1d3eb] via-[#6f46c0] to-[#e0d3ea] rounded-[30px] justify-center items-center inline-flex">
                        <div className="h-[25px] text-center text-black text-[15px] font-medium font-['Inter']">Publier</div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Publication