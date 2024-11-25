import useQueryParams from "@/hooks/useQueryParams";
import { Group } from "@/types/group";
import Image from "next/image";
import Link from "next/link";
import React, { useEffect, useState } from "react";

export const GroupList = () => {
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
            <div className="w-full overflow-scroll no-scrollbar">
                <div className="w-full grid grid-cols-4 gap-5">
                    {groups.map((group, idx) => (
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
                    ))}
                </div>
                <button className="block" onClick={next}>
                    next
                </button>
                <button className="block" onClick={previous}>
                    previous
                </button>
            </div>
        </>
    );
};
