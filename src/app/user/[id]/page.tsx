import { User } from "@/types/user";

export default async function Page({ params }: { params: { id: string } }) {
    const resp = await fetch(`${process.env.API_URL}/api/user/${params.id}`);

    const user: User = await resp.json();

    return <>{JSON.stringify(user)}</>;
}
