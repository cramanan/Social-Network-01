import React from "react";

export const FollowListMobile = () => {
    return (
        <>
            <div className="h-[calc(100vh-111px)] flex flex-col w-full">
                <h2 className="text-4xl text-white text-center py-5">
                    Follow List
                </h2>
            </div>
            <div className="flex flex-col items-center gap-3 mx-5 overflow-scroll no-scrollbar"></div>
        </>
    );
};
