"use client";

import { useState } from "react";
import { Smartphone } from "@/components/icons/smartphone";
import { Register } from "./Register";
import { Login } from "./Login";

export default function Auth() {
    const [isSignUp, setIsSignUp] = useState(false);

    const toggleOverlay = () => setIsSignUp(!isSignUp);

    return (
        <div className="h-full flex items-center justify-center min-h-screen bg-gray-100">
            <div className="flex flex-col md:flex-row w-full max-w-6xl h-full md:h-[748px] relative rounded-[20px] overflow-hidden">
                <div
                    className={`w-full hidden md:flex md:w-1/2 h-1/2 md:h-full relative bg-white/90 shadow border border-white backdrop-blur-[53px] flex items-center justify-center duration-50 transition-transform duration-500 z-10 ${
                        isSignUp && "md:translate-x-full"
                    }`}
                >
                    <Smartphone />
                </div>

                <div
                    className={`h-full pb-40 flex flex-col justify-center items-center w-full md:w-1/2 md:justify-start relative bg-gradient-to-bl from-[#1667e0] to-[#e492e5] shadow border border-white backdrop-blur-[53px] flex flex-col justify-start items-center transition-transform duration-500 ${
                        isSignUp && "md:-translate-x-full"
                    }`}
                >
                    {
                        /*This returns a <form></form>*/
                        isSignUp ? <Register /> : <Login />
                    }
                    <span className="w-4/5 h-[1px] bg-black"></span>
                    <button
                        onClick={toggleOverlay}
                        className="mt-2 text-neutral-100 font-bold m-3"
                    >
                        {isSignUp
                            ? "Already have an account?"
                            : "Create new account"}
                    </button>
                </div>
            </div>
        </div>
    );
}
