"use client";

export default function Home() {
    const onSubmit = (e) => {
        e.preventDefault();

        fetch("/api/register", {
            method: "POST",
            body: JSON.stringify(Object.fromEntries(new FormData(e.target))),
        }).catch(console.error);
    };

    const onSubmitLogin = (e) => {
        e.preventDefault();
        fetch("/api/login", {
            method: "POST",
            body: JSON.stringify(Object.fromEntries(new FormData(e.target))),
        }).catch(console.error);
    };

    return (
        <>
            <form method="POST" onSubmit={onSubmit}>
                <input type="text" name="nickname" placeholder="nickname" />
                <input type="text" name="email" placeholder="email" />
                <input type="password" name="password" placeholder="password" />
                <input type="text" name="firstname" placeholder="firstname" />
                <input type="text" name="lastname" placeholder="lastname" />
                <input type="date" name="dateofbirth" />
                <button type="submit">Register</button>
            </form>
            <br />
            <form onSubmit={onSubmitLogin}>
                <input type="text" name="email" placeholder="email" />
                <input type="password" name="password" placeholder="password" />
                <button type="submit">Login</button>
            </form>
        </>
    );
}
