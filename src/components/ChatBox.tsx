import React from 'react'
import { BackIcon } from './icons/BackIcon'
import { SendIcon } from './icons/SendIcon'
import { EmoteIcon } from './icons/EmoteIcon'

const ChatBox = () => {
    return (
        <>
            <div id='chatBox' className='w-[343px] h-[642px] relative bg-[#fbfbfb] rounded-[25px] flex flex-col'>
                <div className='w-[343px] h-[50px] relative bg-[#445ab3]/20 rounded-tl-[25px] rounded-t-[25px] border-b border-black flex flex-row items-center p-2'>
                    <BackIcon />
                </div>
                <div className='flex flex-col flex-grow overflow-scroll no-scrollbar'>
                    Chat
                </div>
                <div className='h-[50px] flex flex-row justify-between items-center m-5 bg-[#445ab3]/20 rounded-[25px] p-2'>
                    <div className='flex flex-row gap-2'>
                        <EmoteIcon />
                        <input type='text' placeholder='Enter your message' className='bg-white/0'></input>
                    </div>
                    <SendIcon />
                </div>
            </div>
        </>
    )
}

export default ChatBox