export default function Home() {
    return (
        <>
            <form method="POST" action="/api/register">
                <input type="text" name="nickname" placeholder="nickname" />
                <input type="text" name="email" placeholder="email" />
                <input type="password" name="password" placeholder="password" />
                <input type="text" name="firstname" placeholder="firstname" />
                <input type="text" name="lastname" placeholder="lastname" />
                <input type="date" name="dateofbirth" />
            </form>
        </>
    );
}
