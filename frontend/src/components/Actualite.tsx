"use client";

import React, { useEffect, useState } from "react";

// import { Post } from "@/types/post";
// import { QueryParams } from "@/types/query";

import { NewPost } from "./NewPost";
import PostComponent from "./PostComponent";
import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";
import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";
import { Post } from "@/types/post";
import useQueryParams from "@/hooks/useQueryParams";
import Media from "./Media";
import { PostMedia } from "./PostMedia";
// import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";
// import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";

const Actualite = () => {
    const [currentFilter, setCurrentFilter] = useState("All");
    const navStyle =
        "text-black/50 text-xl font-extralight font-['Inter'] tracking-wide";

    const [posts, setPosts] = useState<Post[]>([]);

    const { limit, offset, next, previous } = useQueryParams();

    useEffect(() => {
        fetch(`/api/group/00000000/posts?limit=${limit}&offset=${offset}`)
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setPosts)
            .catch(console.error);
    }, [limit, offset]);

    return (
        <>
            <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                <div className="shadow-xl w-full mb-5 ">
                    <nav
                        className="flex flex-wrap items-center justify-center sm:flex-row sm:justify-between"
                        aria-label="post filter"
                    >
                        <ul className="flex flex-row gap-10 m-4 mt-3 ">
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
                        <div className="flex flex-row">
                            <NewPost />
                        </div>
                    </nav>
                </div>

                <section
                    className="flex flex-col w-full gap-3 overflow-scroll no-scrollbar xl:px-5"
                    aria-label="Posts"
                >
                    {posts.map((post, idx) => (
                        post.images.length > 0 ? (
                            <PostMedia key={idx} post={post} />
                        ) : (
                            <PostComponent key={idx} post={post} />
                        )
                    ))}

                </section>

                {/* <section className="grid grid-cols-3 gap-5 overflow-scroll no-scrollbar">
                    {posts.map((post, idx) => (
                        post.images.length > 0 && <Media key={idx} {...post} />
                    ))}
                </section> */}

                <div className="flex gap-5 p-2">
                    <button onClick={previous}>
                        <CloseSideMenuIcon />
                    </button>
                    <button onClick={next}>
                        <OpenSideMenuIcon />
                    </button>
                </div>
            </div>
        </>
    );
};

export default Actualite;
