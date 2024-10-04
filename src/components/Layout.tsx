"use client";

import React ,{ ReactNode } from "react";
import  Header  from "@/components/Header";
import Chat from "@/components/Chat";

export default function Layout({ children }: { children: ReactNode }) {
    return (
        <>
        <Header></Header>
        <Chat></Chat>
            {children}
            <footer></footer>
        </>
    );
}
