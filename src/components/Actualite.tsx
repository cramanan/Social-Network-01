"use client";

import React, { useEffect, useState } from "react";

// import { Post } from "@/types/post";
// import { QueryParams } from "@/types/query";

import { NewPost } from "./newPost";
import PostComponent from "./PostComponent";
import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";
import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";
import { QueryParams } from "@/types/query";
import { Post } from "@/types/post";
// import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";
// import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";

const Actualite = () => {
    const [currentFilter, setCurrentFilter] = useState("All");
    const navStyle =
        "text-black/50 text-xl font-extralight font-['Inter'] tracking-wide";

    const [posts, setPosts] = useState<Post[]>([]);

    const [params, setParams] = useState<QueryParams>({ limit: 10, offset: 0 });

    const changePage = (n: number) => () => {
        if (params.offset + n * params.limit < 0) return;
        setParams({ ...params, offset: n * 10 });
    };
    useEffect(() => {
        fetch(
            `/api/group/00000000/posts?limit=${params.limit}&offset=${params.offset}`
        )
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setPosts)
            .catch(console.error);
    }, [params.limit, params.offset]);

    return (
        <>
            <div className="flex flex-col relative items-center w-screen h-[calc(100vh-128px)] xl:bg-white/25 z-10 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-60px)]">
                <div className="shadow-xl w-full mb-10 ">
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
                    className="flex flex-col gap-3 mx-3 overflow-scroll no-scrollbar"
                    aria-label="Posts"
                >
                    {posts.map((post, idx) => (
                        <PostComponent post={post} key={idx} />
                    ))}
                </section>
                <div className="flex">
                    <button onClick={changePage(-1)}>
                        <CloseSideMenuIcon />
                    </button>
                    <button onClick={changePage(+1)}>
                        <OpenSideMenuIcon />
                    </button>
                </div>
            </div>
        </>
    );
};

export default Actualite;
