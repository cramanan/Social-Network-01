import Header from "@/components/Header";
import SideNavMenu from "@/components/SideNavMenu";
import React from "react";
import Chat from "@/components/Chat";
import "tailwindcss/tailwind.css";

interface Props {
  children: React.ReactNode;
}

const HomeProfileLayout: React.FC<Props> = ({ children }) => {
  return (
    <div className="flex flex-col h-screen">
      <header className="top-0 w-full">
        <Header />
      </header>
      <div className="flex flex-row flex-grow overflow-hidden">
        <section className="flex flex-1 items-center justify-center">
          <SideNavMenu />
        </section>
        <section className="overflow-auto flex flex-col justify-center items-center flex-grow h-full">
          {children}
        </section>
        <section className="flex justify-end flex-1">
          <Chat />
        </section>
      </div>
    </div>
  );
};

export default HomeProfileLayout;
