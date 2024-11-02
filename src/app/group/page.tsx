"use client";

import { Group } from "@/types/group";
import { QueryParams } from "@/types/query";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function Page() {
    const [groups, setGroups] = useState<Group[]>([]);
    const [params] = useState<QueryParams>({ limit: 10, offset: 0 });

    // const changePage = (n: number) => () => {
    //     if (params.offset + n * params.limit < 0) return;
    //     setParams({ ...params, offset: n * 10 });
    // };

    useEffect(() => {
        fetch(`/api/groups?limit=${params.limit}&offset=${params.offset}`)
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setGroups)
            .catch(console.error); // TODO: edit Global to a valid URL value
    }, [params.limit, params.offset]);

    return (
        <div>
            {groups.map((group, idx) => (
                <Link key={idx} href={`/group/${group.id}`}>
                    {group.name}
                </Link>
            ))}
        </div>
    );
}
