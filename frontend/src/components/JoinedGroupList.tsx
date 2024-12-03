import useQueryParams from "@/hooks/useQueryParams";
import { Group } from "@/types/group";
import Image from "next/image";
import Link from "next/link";
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
                {groups.length > 0 ? (groups.map((group, idx) => (
                    <Link
                        key={idx}
                        href={`/group/${group.id}`}
                        className="flex flex-col justify-center items-center"
                    >
                        <Image
                            src={group.image}
                            alt=""
                            width={56}
                            height={56}
                        />
                        <span>{group.name}</span>
                    </Link>
                )))
                    :
                    <p className="text-center font-bold">
                        No joined group(s) found.
                    </p>
                }

            </ul>
        </>
    )
}
