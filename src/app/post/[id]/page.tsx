import CommentForm from "@/components/CommentForm";
import { Post } from "@/types/post";
import { Params } from "@/types/query";

export default async function Page({ params }: { params: Params }) {
    const { id } = await params;
    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/${id}`
    );

    const post: Post = await response.json();

    return (
        <>
            <div className="whitespace-pre-wrap">
                {JSON.stringify(post, null, "\t")}
            </div>
            <CommentForm post={post} />
        </>
    );
}
