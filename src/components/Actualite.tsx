"use client";

import React, { useEffect, useState } from "react";
import CreatePost from "./CreatePost";
import { Post } from "@/types/post";
import { QueryParams } from "@/types/query";
import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";
import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";

const Actualite = () => {
    const [currentFilter, setCurrentFilter] = useState("All");
    const navStyle =
        "text-black/50 text-xl font-extralight font-['Inter'] tracking-wide";

    const [posts, setPosts] = useState<Post[]>([]);
    console.log(posts);

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
                ></section>
                <div>
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
