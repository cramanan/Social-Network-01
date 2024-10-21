import React from "react";

const ProfileBanner = () => {
    return (
        <div className="flex flex-row items-center h-36">
            <div className="w-36 h-36 z-10 bg-white rounded-full -m-8"></div>
            <div className="w-[440px] h-15 bg-white rounded-r-[30px] flex flex-col justify-between py-1 pl-10">
                <div className="text-black text-2xl font-semibold font-['Inter']">
                    User
                </div>
                <div className="text-black/70 text-xl font-light font-['Inter']">
                    @nickname
                </div>
            </div>
        </div>
    );
};

export default ProfileBanner;