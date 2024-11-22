import FollowButton from "@/components/FollowButton";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { Params } from "@/types/query";
import { User } from "@/types/user";

export default async function Page({ params }: { params: Params }) {
    const { id } = await params;

    const resp = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/user/${id}`
    );

    const user: User = await resp.json();

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col justify-center items-end my-3 mt-11">
                    <ProfileBanner {...user} />
                    <ProfileStats userId={user.id} />
                </div>
                <FollowButton userId={user.id} username={user.nickname} />
                <div className="whitespace-pre-wrap">
                    {JSON.stringify(user, null, "\t")}
                </div>
            </HomeProfileLayout>
        </>
    );
}
