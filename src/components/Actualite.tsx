"use client";

import React, { useState } from "react";

import Post from "./Post";
import { NewPost } from "./icons/newPost";

const Actualite = () => {
    const [currentFilter, setCurrentFilter] = useState("All");
    const navStyle =
        "text-black/50 text-xl font-extralight font-['Inter'] tracking-wide";
    return (
        <>
            <div className="mt-3 flex flex-col items-center w-screen h-[calc(100vh-60px)]  bg-white/25 md:rounded-t-[25px] lg:w-[900px] lg:rounded-t-[25px]">
                <div className=" shadow-xl flex flex-row w-full mb-10 ">
                    <nav aria-label="post filter">
                        <ul className="w-full flex flex-row gap-10 m-4 mt-3 ">
                            {["All", "Publication", "Media"].map((filter) => (
                                <li key={filter} className={navStyle}>
                                    <a
                                        href={`#${filter}`}
                                        onClick={() => setCurrentFilter(filter)}
                                        aria-current={
                                            currentFilter === filter
                                                ? "page"
                                                : undefined
                                        }
                                    >
                                        {filter}
                                    </a>
                                </li>
                            ))}
                        </ul>
                    </nav>
                    <div className="flex w-full justify-end items-end">
                        <NewPost />
                    </div>
                </div>

                <section
                    className="flex flex-col gap-3 mx-3 overflow-scroll no-scrollbar"
                    aria-label="Posts"
                >
                    <Post />
                    <Post />
                    <Post />
                </section>
            </div>
        </>
    );
};

export default Actualite;
