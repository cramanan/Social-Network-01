"use client";

import { useParams } from "next/navigation";
import React from "react";

export default function Page() {
    const { id } = useParams();
    const socket = new WebSocket(
        `ws://${process.env.NEXT_PUBLIC_API_URL}/api/group/${id}/chatroom`
    );
    console.log(id);

    return <div>Page</div>;
}
