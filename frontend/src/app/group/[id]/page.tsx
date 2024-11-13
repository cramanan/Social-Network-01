import React from "react";
import { Group } from "@/types/group";
import { Params } from "@/types/query";
import NewEvent from "./NewEvent";
import Events from "./Events";

export default async function GroupPage({ params }: { params: Params }) {
    const { id } = await params;

    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/group/${id}`
    );
    const group: Group = await response.json();

    return (
        <>
            <h1>{group.name}</h1>
            <p>{group.description}</p>
            <NewEvent groupId={group.id} />
            <div>Events</div>
            <Events groupId={group.id} />
        </>
    );
}
