"use client";
import { useAuth } from "@/hooks/useAuth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { string, z, ZodType } from "zod";
import { useRouter } from "next/navigation";

type LoginFields = {
    email: string;
    password: string;
};

const LoginSchema: ZodType<LoginFields> = z.object({
    email: string()
        .min(1, "Email is required")
        .email("Invalid email format")
        .max(100, "Email is too long"),
    password: string()
        .min(6, "Password must be at least 6 characters")
        .max(50, "Password is too long"),
});

export const Login = () => {
    const router = useRouter();
    const {
        register,
        handleSubmit,
        setError,
        formState: { errors, isSubmitting },
    } = useForm<LoginFields>({
        resolver: zodResolver(LoginSchema),
        mode: "onChange",
    });

    const { login } = useAuth();

    const onSubmit = async ({ email, password }: LoginFields) => {
        try {
            await login(email, password);
            router.push("/");
        } catch {
            setError("root", { message: "An error occurred during login" });
        }
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col h-full justify-center align-center md:justify-center gap-20 md:gap-12 pt-14">
                <h1 className="text-white text-4xl font-semibold font-['Noto Sans']">
                    Login
                </h1>
                <div className="flex flex-col justify-center items-center md:gap-4 py-7">
                    <div className="w-full flex flex-col items-center">
                        <input
                            type="email"
                            className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-2 placeholder-white"
                            {...register("email")}
                            placeholder="Email"
                            aria-label="Email"
                        />
                        {errors.email && (
                            <span className="text-red-500 text-sm mb-2">
                                {errors.email.message}
                            </span>
                        )}
                    </div>

                    {/* Champ Password */}
                    <div className="w-full flex flex-col items-center">
                        <input
                            type="password"
                            {...register("password")}
                            className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-2 placeholder-white"
                            placeholder="Password"
                            aria-label="Password"
                        />
                        {errors.password && (
                            <span className="text-red-500 text-sm mb-2">
                                {errors.password.message}
                            </span>
                        )}
                    </div>
                </div>

                <div className="flex flex-col items-center w-full">
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="w-2/4 bg-white mb-3 hover:bg-violet-100 text-black border-r border-l border-black font-bold py-2 px-4 rounded-md disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                        {isSubmitting ? "Logging in..." : "Log in"}
                    </button>

                    {errors.root && (
                        <div className="text-red-500 text-sm mt-2">
                            {errors.root.message}
                        </div>
                    )}

                    <div className="text-gray-400 text-xs mt-4 text-center">
                        <p>Password requirements:</p>
                        <ul className="list-disc text-left pl-4">
                            <li>At least 6 characters long</li>
                        </ul>
                    </div>
                </div>
            </div>
        </form>
    );
};
