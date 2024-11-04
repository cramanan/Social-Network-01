"use client";

import { Post } from "@/types/post";
import React, { ChangeEvent, FormEvent, useState } from "react";

export default function CommentForm({ post }: { post: Post }) {
    const [content, setContent] = useState("");

    const onChange = (e: ChangeEvent<HTMLTextAreaElement>) =>
        setContent(e.target.value);

    const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const data = new FormData(e.currentTarget);
        data.append("postId", post.id);
        data.append("content", content);

        const response = await fetch(`/api/post/${post.id}/comment`, {
            method: "POST",
            body: data,
        });

        if (response.ok) setContent("");
    };

    return (
        <form onSubmit={onSubmit}>
            <h2>Comment {post.username}&apos;s post:</h2>
            <textarea
                onChange={onChange}
                className="resize-none"
                value={content}
            />
            <div>
                <input type="file" name="images" id="images" />
            </div>
            <button type="submit">Send</button>
        </form>
    );
}
