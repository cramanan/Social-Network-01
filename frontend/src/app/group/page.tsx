"use client";

import { GroupList } from "@/components/GroupList";
import NewGroup from "./NewGroup";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useState } from "react";
import { JoinedGroupList } from "@/components/JoinedGroupList";

export default function Page() {
    const [windows, setWindows] = useState([true, false]);
    const [currentFilter, setCurrentFilter] = useState("Joined groups");

    const changeWindow = (i: number, n: string) => () => {
        setWindows((prev) => prev.map(({ }, idx) => idx === i));
        setCurrentFilter(n)
    }

    const titles = ["Joined groups", "All groups"];
    const bodies = [JoinedGroupList, GroupList];

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <NewGroup />
                    <ul className="flex flex-row w-full justify-evenly">
                        {titles.map((name, idx) => (
                            <li
                                key={idx}
                                className={`w-full text-xl font-['Inter'] text-center cursor-pointer p-3 ${currentFilter === name && "font-bold"} hover:bg-black/10`}
                                onClick={changeWindow(idx, name)}
                            >
                                {name}
                            </li>
                        ))}
                    </ul>
                    {bodies.map(
                        (Value, idx) => windows[idx] && <Value key={idx} />
                    )}
                </div>
            </HomeProfileLayout>
        </>
    );
}
