"use client";

import { useState } from "react";
import { Smartphone } from "@/components/icons/smartphone";
import { Register } from "@/components/Register";
import { Login } from "@/components/Login";

export default function Auth() {
    const [isSignUp, setIsSignUp] = useState(false);

    const toggleOverlay = () => {
        setIsSignUp(!isSignUp);
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="flex flex-col md:flex-row w-full max-w-6xl h-full md:h-[748px] relative rounded-[20px] overflow-hidden">
                <div
                    className={`w-full md:w-1/2 h-1/2 md:h-full relative bg-white/90 shadow border border-white backdrop-blur-[53px] flex items-center justify-center duration-50 transition-transform duration-500 z-10 ${
                        isSignUp && "md:translate-x-full "
                    }`}
                >
                    <Smartphone />
                </div>

                <div
                    className={`w-full md:w-1/2 h-1/2 md:h-full relative bg-gradient-to-bl from-[#1667e0] to-[#e492e5] shadow border border-white backdrop-blur-[53px] flex flex-col justify-start items-center transition-transform duration-500 ${
                        isSignUp && "md:translate-x-[-100%]"
                    }`}
                >
                    {
                        /*This returns a <form></form>*/
                        isSignUp ? <Register /> : <Login />
                    }
                    <button
                        onClick={toggleOverlay}
                        className="w-full md:w-[200px] h-[46px] hover:bg-gradient-to-tr border border-bg-gradient-to-tr from-[#4821f9] via-[#6f46c0] to-[#e0d3ea] rounded-xl flex justify-center items-center text-white text-xl font-semibold font-['Noto Sans']"
                    >
                        {isSignUp ? "Sign in" : "Sign up"}
                    </button>
                </div>
            </div>
        </div>
    );
}
