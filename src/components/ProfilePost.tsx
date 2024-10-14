import React from "react";
import { Like } from "./icons/Like";
import { NewComment } from "./icons/NewComment";
import Comment from "./Comment";

const ProfilePost = () => {
    return (
        <div className="flex flex-row">
            <div className="w-[800px] h-[300px] bg-white rounded-l-[30px] flex flex-col justify-between">
                <textarea
                    className="resize-none w-full h-44 py-4 px-7 rounded-tl-[30px]"
                    defaultValue={
                        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla lectus enim, dignissim id consectetur ut, congue sit amet libero. Nullam in lorem mollis, sollicitudin est et, ornare augue. Suspendisse risus est, porttitor vitae orci eget, sagittis interdum est. Nunc turpis nisl, vestibulum non condimentum eget, eleifend gravida justo."
                    }
                ></textarea>
                <div className="flex flex-row justify-between p-7 pb-5">
                    <div className="text-black/50">
                        Friday 6 september 16:03
                    </div>
                    <div className="flex flex-row gap-20">
                        <Like />
                        <NewComment />
                    </div>
                </div>
            </div>
            <div className="w-[300px] h-[300px] bg-white/20 rounded-r-[30px] flex flex-col justify-center pl-2 pr-10">
                <div className="h-[270px] overflow-scroll no-scrollbar">
                    <Comment />
                    <Comment />
                    <Comment />
                    <Comment />
                    <Comment />
                </div>
            </div>
        </div>
    );
};

export default ProfilePost;
