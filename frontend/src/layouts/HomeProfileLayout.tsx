"use client";

import React, { PropsWithChildren } from "react";
import Header from "@/components/Header";
import SideNavMenu from "@/components/SideNavMenu";
import Chat from "@/components/Chat";
import MobileBottomNav from "@/components/MobileBottomNav";

const HomeProfileLayout = ({ children }: PropsWithChildren) => {
    return (
        <>
            <Header />
            <div className="flex flex-row h-full">
                <div className="hidden items-center xl:flex xl:mt-3">
                    <SideNavMenu />
                </div>

                <main className="flex flex-grow">
                    <div className="w-full h-full relative z-1 xl:absolute xl:left-1/2 xl:-translate-x-1/2 xl:w-fit">
                        {children}
                    </div>
                </main>

                <div className="hidden h-fit right-0 relative z-9 xl:flex xl:mt-3">
                    <Chat />
                </div>
            </div>
            <div className="xl:hidden">
                <MobileBottomNav />
            </div>
        </>
    );
};

export default HomeProfileLayout;
