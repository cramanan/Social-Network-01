import React from "react";

export default function GroupPage({ params }: { params: { id: string } }) {
    return <div>{params.id}</div>;
}
