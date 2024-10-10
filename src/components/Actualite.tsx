import React from 'react'
import CreatePost from './CreatePost'
import Post from './Post'

const Actualite = () => {
  const navStyle = "text-black/50 text-xl font-extralight font-['Inter'] tracking-wide"
  return (
    <div className='mt-3 flex flex-col items-center w-screen h-[calc(100vh-60px)]  bg-white/25 md:w-[600px] md:rounded-t-[25px] lg:w-[900px] lg:rounded-t-[25px]'>
      <div className='mt-3'><CreatePost /></div>
      <ul className='flex flex-row gap-5 mb-5 mt-3'>
        <li className={navStyle} >All</li>
        <li className={navStyle}>Publication</li>
        <li className={navStyle}>Media</li>
      </ul>
      <div className='flex flex-col gap-3 mx-3 overflow-scroll no-scrollbar'>
        <Post />
        <Post />
        <Post />
      </div>
    </div>
  )
}

export default Actualite