"use client";

import { Group } from "@/types/group";
import { QueryParams } from "@/types/query";
import { useEffect, useState } from "react";

export default function GroupPage() {
    const [groups, setGroups] = useState<Group[]>([]);
    const [params, setParams] = useState<QueryParams>({ limit: 10, offset: 0 });

    const changePage = (n: number) => () => {
        if (params.offset + n * params.limit < 0) return;
        setParams({ ...params, offset: n * 10 });
    };

    useEffect(() => {
        fetch(`/api/group?limit=${params.limit}&offset=${params.offset}`)
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setGroups)
            .catch(console.error); // TODO: edit Global to a valid URL value
    }, [params.limit, params.offset]);

    return <div>Groups : {JSON.stringify(groups)}</div>;
}
