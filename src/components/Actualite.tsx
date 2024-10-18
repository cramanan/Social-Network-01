'use client'

import React, { useState } from 'react'
import Post from './Post'

const Actualite = () => {
  const [currentFilter, setCurrentFilter] = useState("All")
  const navStyle = "text-black/50 text-xl font-extralight font-['Inter'] tracking-wide"
  return (
    <>
      <div className='flex flex-col items-center w-screen h-[calc(100vh-68px)] bg-white/25 md:w-[700px] md:rounded-t-[25px]'>
        <h1 className="sr-only">Post feed</h1>

        <nav aria-label="post filter">
          <ul className='flex flex-row gap-10 mb-5 mt-3'>
            {["All", "Publication", "Media"].map((filter) => (
              <li key={filter} className={navStyle}>
                <a href={`#${filter}`} onClick={() => setCurrentFilter(filter)} aria-current={currentFilter === filter ? 'page' : undefined}>
                  {filter}
                </a>
              </li>
            ))}
          </ul>
        </nav>

        <section className='flex flex-col gap-3 mx-3 overflow-scroll no-scrollbar' aria-label="Posts">
          <Post />
          <Post />
          <Post />
        </section>
      </div>
    </>
  )
}

export default Actualite