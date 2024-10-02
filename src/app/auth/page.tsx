"use client";

import { useForm } from "react-hook-form";

interface LoginFields {
    email: string;
    password: string;
}

interface RegisterFields extends LoginFields {
    nickname: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
}

export default function Auth() {
    const { register, handleSubmit: handleRegister } =
        useForm<RegisterFields>();

    const { register: login, handleSubmit: handleLogin } =
        useForm<LoginFields>();

    const registerSubmit = (data: RegisterFields) => {
        fetch("/api/register", {
            method: "POST",
            body: JSON.stringify(data),
        })
            .then((resp) => resp.json())
            .then(console.log)
            .catch(console.error);
    };

    const registerLogin = (data: LoginFields) => {
        fetch("/api/login", {
            method: "POST",
            body: JSON.stringify(data),
        })
            .then((resp) => resp.json())
            .then(console.log)
            .catch(console.error);
    };

    return (
        <>
            <h1>Login</h1>
            <form onSubmit={handleLogin(registerLogin)}>
                <input type="email" placeholder="email" {...login("email")} />
                <input
                    type="password"
                    placeholder="password"
                    {...login("password")}
                />

                <button type="submit">Send</button>
            </form>

            <h1>Register</h1>
            <form onSubmit={handleRegister(registerSubmit)}>
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
