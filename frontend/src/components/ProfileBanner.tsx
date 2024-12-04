import { User } from "@/types/user";
import Image from "next/image";
import React from "react";
import { SettingIcon } from "./icons/SettingIcon";
import { useAuth } from "@/hooks/useAuth";

const ProfileBanner = ({ id, nickname, firstName, lastName, image }: User) => {
    const { user } = useAuth()
    return (
        <div className="flex flex-row items-center h-20">
            <Image
                width={144}
                height={144}
                alt=""
                src={image}
                className="w-36 h-36 z-10 bg-white rounded-full -m-8"
                priority
            />
            <div className="flex flex-row min-w-[28vw] h-16 bg-white rounded-r-[30px] justify-between items-center py-1 pl-10 pr-3">
                <div className="flex flex-col">
                    <div className="text-black text-2xl font-semibold font-['Inter']">
                        {firstName} {lastName}
                    </div>
                    <div className="text-black/70 text-base font-light font-['Inter']">
                        @{id} / @{nickname}
                    </div>
                </div>
                {id === user?.id && <a href="/profile/settings" className="bg-black/50 rounded-full">
                    <SettingIcon />
                </a>}
            </div>
        </div>
    );
};

export default ProfileBanner;
