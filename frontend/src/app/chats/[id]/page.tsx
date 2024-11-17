import { User } from "@/types/user";
import { Params } from "@/types/query";
import ChatBox from "@/components/ChatBox";

export default async function Page({ params }: { params: Params }) {
    const { id } = await params;

    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/user/${id}`
    );

    const user: User = await response.json();

    return (
        <>
            <ChatBox onClose={() => {}} recipient={user} />
            {/* <h1 className="flex justify-between font-bold p-2">
                <Link href="/chats">
                    <BackIcon />
                </Link>
                <Link
                    href={`/user/${user.id}`}
                    className="flex flex-col items-center"
                >
                    <Image
                        src={user.image}
                        alt={`${user.nickname}'s profile picture`}
                        width={40}
                        height={40}
                        priority
                    />
                    <div>{user.nickname}</div>
                </Link>
                <span />
            </h1>
            <ChatRoom recipient={user} /> */}
        </>
    );
}
