import Link from "next/link";
import ChatRoom from "./ChatRoom";

type Params = Promise<{ id: string }>;

export default async function Page(props: { params: Params }) {
    const params = await props.params;
    const id = params.id;

    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/user/${id}`
    );

    const user = await response.json();

    return (
        <>
            <h1 className="flex justify-between font-bold p-2">
                <Link href="/chats" className="">
                    &lt;-
                </Link>
                <span>{user.nickname}</span>
                <span />
            </h1>
            <ChatRoom recipient={user} />
        </>
    );
}
