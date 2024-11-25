import { User } from "@/types/user";
import { Params } from "@/types/query";
import ChatBox from "@/components/ChatBox";

export default async function Page({ params }: { params: Params }) {
    const { id } = await params;

    const response = await fetch(
        `http://${process.env.NEXT_PUBLIC_API_URL}/api/user/${id}`
    );

    const user: User = await response.json();

    return (
        <>
            <ChatBox recipient={user} />
        </>
    );
}
