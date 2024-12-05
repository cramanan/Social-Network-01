"use client";

import { useAuth } from "@/hooks/useAuth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { z, ZodType } from "zod";

// Form Datas
type RegisterFields = {
    email: string;
    password: string;
    nickname: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
};

// Zod Schema for the resolver
const RegisterSchema: ZodType<RegisterFields> = z.object({
    email: z.string().email(),
    password: z
        .string()
        .min(8, { message: "Password is too short" })
        .max(20, { message: "Password is too long" }),
    nickname: z.string(),
    firstName: z.string(),
    lastName: z.string(),
    dateOfBirth: z.string().date("Invalid Date"),
});

export const Register = () => {
    const router = useRouter();
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<RegisterFields>({
        resolver: zodResolver(RegisterSchema),
    });
    const { signup } = useAuth();

    const onSubmit = async ({
        nickname,
        email,
        password,
        firstName,
        lastName,
        dateOfBirth,
    }: RegisterFields) => {
        try {
            await signup(
                nickname,
                email,
                password,
                firstName,
                lastName,
                dateOfBirth
            );
            router.push("/");
        } catch {}
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col  w-full gap-20 md:gap-12 pt-14">
                <h1 className="text-white  text-4xl font-semibold font-['Noto Sans']">
                    Register
                </h1>
                <div className="flex flex-col justify-center items-center">
                    <input
                        type="text"
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        {...register("nickname")}
                        placeholder="Nickname"
                        aria-label="Nickname"
                    />
                    <input
                        type="email"
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        {...register("email")}
                        placeholder="Email"
                        aria-label="Email"
                    />
                    {errors.email && (
                        <span className="text-red-500 text-sm mb-2">
                            {errors.email.message}
                        </span>
                    )}
                    <input
                        type="password"
                        {...register("password")}
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        placeholder="Password"
                        aria-label="Password"
                    />
                    {errors.password && (
                        <span className="text-red-500 text-sm mb-2">
                            {errors.password.message}
                        </span>
                    )}
                    <input
                        type="text"
                        {...register("firstName")}
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        placeholder="First Name"
                        aria-label="First Name"
                    />

                    <input
                        type="text"
                        {...register("lastName")}
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        placeholder="Last Name"
                        aria-label="Last Name"
                    />

                    <input
                        type="date"
                        {...register("dateOfBirth")}
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        placeholder="Date of birth"
                        aria-label="Date of birth"
                    />
                    {errors.dateOfBirth && (
                        <span className="text-red-500 text-sm mb-2">
                            {errors.dateOfBirth.message}
                        </span>
                    )}
                </div>
                <div className="flex justify-center w-full">
                    <button
                        disabled={isSubmitting}
                        type="submit"
                        className="w-2/4 bg-white mb-4 hover:bg-violet-100 text-black border-r border-l border-black font-bold py-2 px-4 rounded-md"
                    >
                        {isSubmitting ? "Signing in..." : "Sign in"}
                    </button>
                </div>
            </div>
        </form>
    );
};
