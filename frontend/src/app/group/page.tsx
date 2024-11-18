"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { Group } from "@/types/group";
import Link from "next/link";
import { useEffect, useState } from "react";
import NewGroup from "./NewGroup";
import Image from "next/image";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default function Page() {
    const [groups, setGroups] = useState<Group[]>([]);
    const { limit, offset, next, previous } = useQueryParams();

    useEffect(() => {
        fetch(`/api/groups?limit=${limit}&offset=${offset}`)
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setGroups)
            .catch(console.error); // TODO: edit Global to a valid URL value
    }, [limit, offset]);

    return (
        <>
            <HomeProfileLayout >
                <div className="overflow-scroll h-full">
                    <NewGroup />
                    <div>
                        {groups.map((group, idx) => (
                            <Link key={idx} href={`/group/${group.id}`}>
                                <Image
                                    src={group.image}
                                    alt=""
                                    width={56}
                                    height={56}
                                    className="w-14 h-14"
                                />
                                <span>{group.name}</span>
                            </Link>
                        ))}
                        <button className="block" onClick={next}>
                            next
                        </button>
                        <button className="block" onClick={previous}>
                            previous
                        </button>
                    </div>
                </div>
            </HomeProfileLayout>
        </>
    );
}
