import { useAuth } from "@/providers/AuthContext";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { z, ZodType } from "zod";

// Form Datas
interface RegisterFields {
    email: string;
    password: string;
    nickname: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
}

// Zod Schema for the resolver
export const UserSchema: ZodType<RegisterFields> = z.object({
    email: z.string().email(),
    password: z
        .string()
        .min(8, { message: "Password is too short" })
        .max(20, { message: "Password is too long" }),
    nickname: z.string(),
    firstName: z.string(),
    lastName: z.string(),
    dateOfBirth: z.string().date("invalid date"),
});

export const Register = () => {
    const { register, handleSubmit } = useForm<RegisterFields>({
        resolver: zodResolver(UserSchema),
    });

    const { setUser } = useAuth();

    const router = useRouter();

    const onSubmit = (data: RegisterFields) => {
        fetch("/api/register", {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then((resp) => {
                if (resp.ok) return resp.json();
                throw "Error";
            })
            .then(setUser)
            .then(() => router.push("/"))

            .catch(console.error);
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col  w-full gap-20 md:gap-12 p-14">
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
                    <input
                        type="password"
                        {...register("password")}
                        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
                        placeholder="Password"
                        aria-label="Password"
                    />
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
                </div>
                <button type="submit">Sign up</button>
            </div>
        </form>
    );
};
