import Image from "next/image";
import React from "react";

const ProfileBanner = ({
    image,
    id,
    firstName,
}: {
    image: string;
    id: string;
    firstName: string;
}) => {
    return (
        <div className="flex flex-row items-center h-36">
            <Image
                width={144}
                height={144}
                alt=""
                src={image}
                className="w-36 h-36 z-10 bg-white rounded-full -m-8"
            ></Image>
            <div className="w-[440px] h-15 bg-white rounded-r-[30px] flex flex-col justify-between py-1 pl-10">
                <div className="text-black text-2xl font-semibold font-['Inter']">
                    {firstName}
                </div>
                <div className="text-black/70 text-xl font-light font-['Inter']">
                    @{id}
                </div>
            </div>
        </div>
    );
};

export default ProfileBanner;
