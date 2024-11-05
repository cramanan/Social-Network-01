import Header from "@/components/Header";
import SideNavMenu from "@/components/SideNavMenu";
import React from "react";
import Chat from "@/components/Chat";
import MobileBottomNav from "@/components/mobile/MobileBottomNav";

interface Props {
    children: React.ReactNode;
}

const HomeProfileLayout: React.FC<Props> = ({ children }) => {
    return (
        <>
            <Header />

            <div className="flex flex-row h-full">
                <div className="hidden items-center left-0 top-[150px] xl:flex xl:mt-3">
                    <SideNavMenu />
                </div>

                <main className="flex flex-grow">
                    {/* <div className="absolute left-1/2 -translate-x-1/2 xl:mt-3"> */}
                    {children}
                    {/* </div> */}
                </main>

                <div className="hidden right-0 xl:flex xl:mt-3">
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
