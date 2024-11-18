import React, { useState } from "react";
import { LikeIcon } from "./icons/LikeIcon";

const Comment = () => {
    const [isLiked, setIsLiked] = useState(false)

    const handleLikeCLick = () => setIsLiked(!isLiked)
    return (
        <>
            <div
                className="flex items-center justify-between relative w-full h-[54px] bg-[#f6f6f6]/0"
                aria-label="Comment"
            >
                <div className="flex flex-row items-center">
                    <div className="m-2">
                        <div className="w-[41px] h-10 bg-[#b53695] rounded-[100px]"></div>
                    </div>
                    <div className="flex flex-col">
                        <span className="h-[21px] text-black text-[15px] font-semibold font-['Inter']">
                            Name
                        </span>
                        <span className=" text-black text-[13px] font-normal font-['Inter']">
                            sss
                        </span>
                    </div>
                </div>
                <div onClick={handleLikeCLick} className="mr-2">
                    <LikeIcon isLiked={isLiked} />
                </div>
            </div>
        </>
    );
};

export default Comment;
