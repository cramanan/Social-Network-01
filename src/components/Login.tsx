import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";

// Form Datas
interface LoginFields {
    email: string;
    password: string;
}

export const Login = () => {
    const { register, handleSubmit } = useForm<LoginFields>();

    const router = useRouter();

    const onSubmit = (/*data: LoginFields*/) => {
        // fetch("/api/login", {
        //     method: "POST",
        //     body: JSON.stringify(data),
        // }).then(()=>{
        router.push("/");
        //});
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col  w-full gap-20 md:gap-12 p-14">
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
                <button type="submit">Sign in</button>
            </div>
        </form>
    );
};
