"use client";

import useQueryParams from "@/hooks/useQueryParams";
import Link from "next/link";
import { Group } from "@/types/group";
import { useEffect, useState } from "react";

export default function Page() {
    const [groups, setGroups] = useState<Group[]>([]);
    const { limit, offset, next, previous } = useQueryParams();

    useEffect(() => {
        fetch(`/api/groups?limit=${limit}&offset=${offset}`)
            .then((resp) => (resp.ok ? resp.json() : [])) // TODO: handle error
            .then(setGroups)
            .catch(console.error);
    }, [limit, offset]);

    return (
        <>
            <div className="flex flex-col overflow-auto">
                {groups.map((group, idx) => (
                    <Link key={idx} href={`/group/${group.id}`}>
                        {group.name}
                    </Link>
                ))}
            </div>
            <button className="block" onClick={next}>
                next
            </button>
            <button className="block" onClick={previous}>
                previous
            </button>
        </>
    );
}
