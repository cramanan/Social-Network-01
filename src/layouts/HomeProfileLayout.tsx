import Header from "@/components/Header";
import SideNavMenu from "@/components/SideNavMenu";
import React from "react";
import Chat from "@/components/Chat";

interface Props {
  children: React.ReactNode;
}

const HomeProfileLayout: React.FC<Props> = ({ children }) => {
  return (
    <>
      <Header />

      <div className="hidden absolute left-0 top-[150px] xl:flex">
        <SideNavMenu />
      </div>

      <div className="absolute mt-3  left-1/2 -translate-x-1/2">
        {children}
      </div>

      <div className="hidden absolute right-0 xl:flex">
        <Chat />
      </div>
    </>
  );
};

export default HomeProfileLayout;
