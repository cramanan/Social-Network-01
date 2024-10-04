import React from "react";

export default function Chat() {
    const imageStyle = "w-10 h-10 bg-white"
    return (
        <>
            <div className="w-60 h-10 bg-white bg-opacity-40 m-5 border border-neutral-400 rounded-b-3xl">
                <div className="flex justify-between items-center px-7">
                    <image className={imageStyle} />
                    <image className={imageStyle} />
                    <image className={imageStyle} />
                </div>
            </div>
        </>
    );
}

// width: 343px;
// height: 50px;
// top: 112.18px;
// left: 1553px;
// gap: 0px;
// border-radius: 0px 0px 30px 30px;
// border: 2px 0px 0px 0px;
// opacity: 0px;
// angle: 0.36 deg;
