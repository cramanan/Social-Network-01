import Image from "next/image";
import React from "react";

const ProfileBanner = ({
    image,
    id,
    firstName,
    lastName,
    nickname,
}: {
    image: string;
    id: string;
    firstName: string;
    lastName: string;
    nickname: string;
}) => {
    return (
        <div className="flex flex-row items-center h-36">
            <Image
                width={144}
                height={144}
                alt=""
                src={image}
                className="w-36 h-36 z-10 bg-white rounded-full -m-8"
            />
            <div className="flex flex-col min-w-[28vw] h-16 bg-white rounded-r-[30px] justify-between py-1 pl-10">
                <div className="text-black text-2xl font-semibold font-['Inter']">
                    {firstName} {lastName}
                </div>
                <div className="text-black/70 text-base font-light font-['Inter']">
                    @{nickname}
                </div>
            </div>
        </div>
    );
};

export default ProfileBanner;
