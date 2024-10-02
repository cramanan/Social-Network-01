"use client";

import { useForm } from "react-hook-form";

type RegisterData = {
    nickname: string;
    email: string;
    password: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
};

export default function Auth() {
    const { register, handleSubmit } = useForm<RegisterData>();

    const onSubmit = (data: RegisterData) => {
        fetch("/api/register", {
            method: "POST",
            body: JSON.stringify(data),
        })
            .then((resp) => resp.json())
            .then(console.log)
            .catch(console.error);
    };

    return (
        <>
            <form onSubmit={handleSubmit(onSubmit)}>
                <input
                    type="text"
                    placeholder="nickname"
                    {...register("nickname")}
                />
                <input
                    type="email"
                    placeholder="email"
                    {...register("email")}
                />
                <input
                    type="password"
                    placeholder="password"
                    {...register("password")}
                />
                <input
                    type="date"
                    placeholder="date of birth"
                    {...register("dateOfBirth")}
                />
                <input
                    type="text"
                    placeholder="first name"
                    {...register("firstName")}
                />
                <input
                    type="text"
                    placeholder="last name"
                    {...register("lastName")}
                />

                <button type="submit">Send</button>
            </form>
        </>
    );
}
