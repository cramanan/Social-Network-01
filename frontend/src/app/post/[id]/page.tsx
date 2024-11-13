import { Post } from "@/types/post";
import { Params } from "@/types/query";
import CommentForm from "./CommentForm";
import Publication from "@/components/Publication";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default async function Page({ params }: { params: Params }) {
    const { id } = await params;
    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/${id}`
    );

    const post: Post = await response.json();

    return (
        <>
            <HomeProfileLayout >
                <div className="whitespace-pre-wrap">
                    {JSON.stringify(post, null, "\t")}
                </div>
                <CommentForm post={post} />
            </HomeProfileLayout>
        </>

    );
}
