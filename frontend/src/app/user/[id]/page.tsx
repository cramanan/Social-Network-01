import FollowButton from "@/components/FollowButton";
import ProfileStats from "@/components/ProfileStats";
import { Params } from "@/types/query";
import { User } from "@/types/user";

export default async function Page({ params }: { params: Params }) {
    const { id } = await params;

    const resp = await fetch(
        `http://${process.env.NEXT_PUBLIC_API_URL}/api/user/${id}`
    );

    const user: User = await resp.json();

    return (
        <>
            <div className="whitespace-pre-wrap">
                {JSON.stringify(user, null, "\t")}
            </div>
            <ProfileStats userId={user.id} />
            <FollowButton userId={user.id} username={user.nickname} />
        </>
    );
}
