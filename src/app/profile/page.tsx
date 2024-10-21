import Chat from "@/components/Chat";
import Header from "@/components/Header";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default function Profile() {
    return (
        <div className=" max-w-screen min-h-screen flex flex-col">
            <Header />

            <div className="flex flex-1">
                {/* Contenu de la section gauche */}
                <div className="flex-1 w-full h-full mt-32">
                    <SideNavMenu />
                </div>
                <div className="flex-1 w-full h-full  ">
                    {/* Contenu de la section centrale */}
                    <div className="flex justify-center mb-8">
                        <ProfileBanner />
                    </div>
                </div>
                <div className="flex-1 w-full flex  justify-end ">
                    {/* Contenu de la section droite */}
                    <Chat />
                </div>
            </div>
        </div>
    );
}
