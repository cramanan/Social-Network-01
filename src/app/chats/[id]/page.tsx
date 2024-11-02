import ChatRoom from "./ChatRoom";
import Image from "next/image";
import { BackIcon } from "@/components/icons/BackIcon";
import Link from "next/link";

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
                <Link href="/chats">
                    <BackIcon />
                </Link>
                <div className="flex flex-col items-center">
                    <Image
                        src={user.image}
                        alt={`${user.nickname}'s profile picture`}
                        width={40}
                        height={40}
                        priority
                    />
                    <div>{user.nickname}</div>
                </div>
                <span />
            </h1>
            <ChatRoom recipient={user} />
        </>
    );
}
