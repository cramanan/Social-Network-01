"use client";

import { useAuth } from "@/hooks/useAuth";
import { useForm } from "react-hook-form";

// Form Datas
type LoginFields = {
    email: string;
    password: string;
};

export const Login = () => {
    const { register, handleSubmit } = useForm<LoginFields>();

    const { login } = useAuth();

    const onSubmit = ({ email, password }: LoginFields) =>
        login(email, password);

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col h-full justify-center align-center md:justify-center gap-20 md:gap-12 pt-14">
                <h1 className="text-white  text-4xl font-semibold font-['Noto Sans']">
                    Login
                </h1>
                <div className="flex flex-col justify-center items-center md:gap-4 py-7">
                    <input
                        type="email"
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        {...register("email")}
                        placeholder="Email"
                        aria-label="Email"
                    />
                    <input
                        type="password"
                        {...register("password")}
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        placeholder="Password"
                        aria-label="Password"
                    />
                </div>
                <div className="flex justify-center w-full">
                    <button type="submit" className="w-2/4 bg-white mb-6 hover:bg-violet-100 text-black border-r border-l border-black font-bold py-2 px-4 rounded-md">
                        Log in
                    </button>
                </div>
            </div>
        </form>
    );
};
