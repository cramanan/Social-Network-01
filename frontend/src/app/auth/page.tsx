"use client";

import { useState } from "react";
import { Smartphone } from "@/components/icons/smartphone";
import { Register } from "./Register";
import { Login } from "./Login";

export default function Auth() {
    const [isSignUp, setIsSignUp] = useState(false);

    const toggleOverlay = () => setIsSignUp(!isSignUp);

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="flex flex-col md:flex-row w-full max-w-6xl h-full md:h-[748px] relative rounded-[20px] overflow-hidden">
                <div
                    className={`w-full md:w-1/2 h-1/2 md:h-full relative bg-white/90 shadow border border-white backdrop-blur-[53px] items-center justify-center duration-50 transition-transform duration-500 z-10 ${isSignUp && "md:translate-x-full"
                        }`}
                >
                    <Smartphone />
                </div>

                <div
                    className={`w-full md:w-1/2 md:h-full relative bg-gradient-to-bl from-[#1667e0] to-[#e492e5] shadow border border-white backdrop-blur-[53px] flex flex-col justify-center items-center transition-transform duration-500 ${isSignUp && "md:-translate-x-full"
                        }`}
                >
                    {
                        /*This returns a <form></form>*/
                        isSignUp ? <Register /> : <Login />
                    }
                    <span className="w-4/5 h-1 bg-black"></span>
                    <button onClick={toggleOverlay} className="m-3">
                        {isSignUp ? "Sign in" : "Sign up"}
                    </button>
                </div>
            </div>
        </div>
    );
}
