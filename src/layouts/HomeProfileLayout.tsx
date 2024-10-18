import Chat from "@/components/Chat";
import Header from "@/components/Header";
import SideNavMenu from "@/components/SideNavMenu";
import React, { FC } from "react";

interface Props {
  children: React.ReactNode;
}

const HomeProfileLayout: FC<Props> = ({ children }) => {
  return (
    <div className="flex flex-row min-h-screen">
      <header className="fixed w-full">
        <Header />
      </header>
      <div className="flex flex-1 flex-row w-full mt-20">
        <div className="mt-24 flex flex-1 w-1/3">
          <SideNavMenu />
        </div>
        <main className="flex flex-1 w-1/3 justify-center items-center">
          <div className="w-full h-full">{children}</div>
        </main>
        <div className="flex flex-1 justify-end w-1/3">
          <Chat />
        </div>
      </div>
    </div>
  );
};

export default HomeProfileLayout;
