import { User } from "@/types/user";
import { notFound } from "next/navigation";

export default async function Page({ params }: { params: { id: string } }) {
    let user: User;
    try {
        user = await fetch(`${process.env.API_URL}/api/user/${params.id}`).then(
            (resp) => {
                if (!resp.ok) throw { error: "User not found or private" };
                return resp.json();
            }
        );
    } catch (error) {
        notFound();
    }

    return <>{JSON.stringify(user)}</>;
}
