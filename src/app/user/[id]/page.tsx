import FollowButton from "@/components/FollowButton";
import ProfileStats from "@/components/ProfileStats";
import { User } from "@/types/user";

export default async function Page({ params }: { params: { id: string } }) {
    const resp = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/user/${params.id}`
    );

    const user: User = await resp.json();

    return (
        <>
            <div>{JSON.stringify(user)}</div>
            <ProfileStats userId={user.id} />
            <FollowButton userId={user.id} username={user.nickname} />
        </>
    );
}
