'use client'

import React, { useState } from 'react'
import { ProfileCircle } from './icons/ProfileCircle'

const CreatePost = () => {
    const [postContent, setPostContent] = useState('')
    const [error, setError] = useState('')

    const handleSubmit = (e) => {
        e.preventDefault()
        if (postContent.trim() === '') {
            setError('Please write some content')
        } else {
            //send to db
            console.log("Submitting ...");
            setPostContent('')
            setError('')
        }
    }

    return (
        <>
            <form onSubmit={handleSubmit} className='flex flex-row items-center h-[70px] relative bg-white/95 rounded-[30px] justify-between' aria-labelledby="create-post-title">
                <h2 id="create-post-title" className="sr-only">Create a new post</h2>

                <div>
                    <div className='flex flex-row gap-5 ml-5'>
                        <div aria-hidden="true"><ProfileCircle /></div>
                        <label htmlFor="post-input" className="sr-only">Post content</label>
                        <input id="post-input" value={postContent} onChange={(e) => setPostContent(e.target.value)} placeholder='Create your post' className='md:w-[250px] bg-white/0 outline-none resize-none overflow-scroll no-scrollbar place-content-center' aria-required="true" aria-invalid={error ? "true" : "false"} aria-describedby={error ? "post-error" : undefined}></input>
                    </div>
                </div>

                <button type="submit" className='flex items-center justify-center mr-5 bg-gradient-to-tr from-[#4821f9] via-[#6f46c0] to-[#e0d3ea] rounded-[30px]' aria-label="Submit post">
                    <span className=" text-black text-base font-semibold font-['Inter'] px-3">post</span>
                </button>
            </form>
            {error && (
                <p id="post-error" className="text-red-500 text-center" role="alert">
                    {error}
                </p>
            )}
        </>
    )
}

export default CreatePost