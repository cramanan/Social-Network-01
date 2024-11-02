import { Params } from "@/types/query";
import React from "react";

export default async function GroupPage({ params }: { params: Params }) {
    const { id } = await params;
    return <div>{id}</div>;
}
