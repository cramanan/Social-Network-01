import Actualite from "@/components/Actualite";
import Chat from "@/components/Chat";
import Header from "@/components/Header";

import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import SideNavMenu from "@/components/SideNavMenu";

export default function Profile() {
  return (
    <div className=" max-w-screen min-h-screen flex flex-col">
      <Header />

      <div className="flex flex-1 ">
        {/* Contenu de la section gauche */}

        <div className="flex-1 w-full h-full mt-24">
          <SideNavMenu />
        </div>

        <div className="flex-1 w-full h-full mt-4">
          {/* Contenu de la section centrale */}
          <div className="flex justify-center ">
            <ProfileBanner />
          </div>
          <div className="flex justify-center ml-28 -mt-7 mb-20 ">
            <ProfileStats />
          </div>

          <Actualite />
        </div>
        <div className="flex-1 w-full flex justify-end">
          {/* Contenu de la section droite */}
          <Chat />
        </div>
      </div>
    </div>
  );
}
