"use client";

import { User } from "@/types/user";
import ChatBox from "@/components/ChatBox";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";

export default function Page() {
    const { id } = useParams();
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUser = async () => {
            setLoading(true);
            try {
                const response = await fetch(`/api/user/${id}`);

                const data: User = await response.json();
                setUser(data);
            } catch (error) {
                console.error(error);
            } finally {
                setLoading(false);
            }
        };

        fetchUser();
    }, [id]);

    if (loading) return <>Loading</>;

    if (!user) return <>Could not retrieve user</>;

    return <ChatBox recipient={user} />;
}
