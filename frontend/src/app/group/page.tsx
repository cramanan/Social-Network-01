"use client";

import { GroupList } from "@/components/GroupList";
import NewGroup from "./NewGroup";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useState } from "react";

export default function Page() {
    const [windows, setWindows] = useState([true, false]);

    const changeWindow = (i: number) => () =>
        setWindows((prev) => prev.map(({}, idx) => idx === i));

    const titles = ["Joined groups", "All groups"];
    const bodies = [() => <div>Joined</div>, () => <GroupList />];

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col items-center w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <NewGroup />
                    <ul className="flex flex-row w-full justify-evenly">
                        {titles.map((name, idx) => (
                            <li
                                key={idx}
                                className="w-1/3 text-center cursor-pointer p-3 hover:bg-black/10"
                                onClick={changeWindow(idx)}
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
