"use client";

import React, { useEffect, useState } from "react";
import CreatePost from "./CreatePost";
import { Post } from "@/types/post";
import PostComponent from "./Post";
import { QueryParams } from "@/types/query";
import { BackSideBarLeft } from "./icons/BackSideBarLeft";
import { BackSideBarRight } from "./icons/BackSideBarRight";

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
            `/api/group/Global/posts?limit=${params.limit}&offset=${params.offset}`
        )
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setPosts)
            .catch(console.error); // TODO: edit Global to a valid URL value
    }, [params.limit, params.offset]);

    return (
        <>
            <div className="mt-3 flex flex-col items-center w-screen h-[calc(100vh-60px)]  bg-white/25 md:w-[600px] md:rounded-t-[25px] lg:w-[700px] lg:rounded-t-[25px]">
                <h1 className="sr-only">Post feed</h1>
                <section className="mt-3" aria-label="Create new post">
                    <CreatePost />
                </section>
                <nav aria-label="post filter">
                    <ul className="flex flex-row gap-10 mb-5 mt-3">
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
                <section
                    className="flex flex-col gap-3 mx-3 overflow-scroll no-scrollbar"
                    aria-label="Posts"
                >
                    {posts.map((post, idx) => (
                        <PostComponent key={idx} post={post} />
                    ))}
                </section>
                <div>
                    <button onClick={changePage(-1)}>
                        <BackSideBarLeft />
                    </button>
                    <button onClick={changePage(+1)}>
                        <BackSideBarRight />
                    </button>
                </div>
            </div>
        </>
    );
};

export default Actualite;
