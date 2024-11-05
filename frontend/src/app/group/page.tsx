"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { Group } from "@/types/group";
import Link from "next/link";
import { useEffect, useState } from "react";

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
        <div>
            {groups.map((group, idx) => (
                <Link key={idx} href={`/group/${group.id}`}>
                    {group.name}
                </Link>
            ))}
            <button className="block" onClick={next}>
                next
            </button>
            <button className="block" onClick={previous}>
                previous
            </button>
        </div>
    );
}
