"use client";

import React from "react";
import { ProfileCircle } from "./icons/ProfileCircle";
import { useForm } from "react-hook-form";

type PostRequest = {
    content: string;
};

const CreatePost = () => {
    const {
        register,
        handleSubmit,
        // formState: { errors },
        reset,
    } = useForm<PostRequest>();

    const onSubmit = (data: PostRequest) =>
        fetch("/api/post", {
            method: "POST",
            body: JSON.stringify(data),
        })
            .then((resp) => {
                if (resp.ok) return reset();
                throw "Error";
            })
            .catch(console.error);
    return (
        <>
            <form
                onSubmit={handleSubmit(onSubmit)}
                className="flex flex-row items-center h-[70px] relative bg-white/95 rounded-[30px] w-[500px] justify-between"
                aria-labelledby="create-post-title"
            >
                <h2 id="create-post-title" className="sr-only">
                    Create a new post
                </h2>

                <div>
                    <div className="flex flex-row gap-5 ml-5">
                        <div aria-hidden="true">
                            <ProfileCircle />
                        </div>
                        <label htmlFor="post-input" className="sr-only">
                            Post content
                        </label>
                        <input
                            id="post-input"
                            placeholder="Create your post"
                            className="md:w-[250px] bg-white/0 outline-none resize-none overflow-scroll no-scrollbar place-content-center"
                            aria-required="true"
                            required
                            // aria-invalid={error ? "true" : "false"}
                            // aria-describedby={error ? "post-error" : undefined}
                            {...register("content")}
                        />
                    </div>
                </div>

                <button
                    type="submit"
                    className="flex items-center justify-center mr-5 w-[85px] bg-gradient-to-tr from-[#4821f9] via-[#6f46c0] to-[#e0d3ea] rounded-[30px]"
                    aria-label="Submit post"
                >
                    <span className=" text-white text-base font-semibold font-['Inter']">
                        POST
                    </span>
                </button>
            </form>
            {/* {error && (
                <p
                    id="post-error"
                    className="text-red-500 text-center"
                    role="alert"
                >
                    {error}
                </p>
            )} */}
        </>
    );
};

export default CreatePost;
