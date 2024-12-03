import useQueryParams from "@/hooks/useQueryParams";
import { Group } from "@/types/group";
import React, { useEffect, useState } from 'react'

export const JoinedGroupList = () => {
    const [groups, setGroups] = useState<Group[]>([]);
    const { limit, offset } = useQueryParams();

    useEffect(() => {
        fetch(`/api/profile/groups`)
            .then((resp) => (resp.ok ? resp.json() : []))
            .then(setGroups)
            .catch(console.error); // TODO: edit Global to a valid URL value
    }, [limit, offset]);
    return (
        <>
            <ul>
                {groups.map((group, idx) => (
                    <li key={idx}>{group.name}</li>
                ))}
            </ul>
        </>
    )
}
