import "./page.css";

export default function Docs() {
    const routes = [
        {
            method: "POST",
            route: "/api/register",
            body: {
                nickname: "string",
                email: "string",
                password: "string",
                firstName: "string",
                lastName: "string",
                dateofbirth: "string",
            },
            response: {
                id: "string",
                nickname: "string",
                email: "string",
                firstName: "string",
                lastName: "string",
                dateOfBirth: "string",
                imagePath: "string | null",
                aboutMe: "string | null",
                private: "bool",
                timestamp: "string",
            },
        },
        {
            method: "POST",
            route: "/api/login",
            body: {
                email: "string",
                password: "string",
            },
            response: {
                id: "string",
                nickname: "string",
                email: "string",
                firstName: "string",
                lastName: "string",
                dateOfBirth: "string",
                imagePath: "string | null",
                aboutMe: "string | null",
                private: "bool",
                timestamp: "string",
            },
        },
        {
            method: "GET",
            route: "/api/user/{id}",
            response: {
                id: "string",
                nickname: "string",
                email: "string",
                firstName: "string",
                lastName: "string",
                dateOfBirth: "string",
                imagePath: "string | null",
                aboutMe: "string | null",
                private: "bool",
                timestamp: "string",
            },
        },
        {
            method: "POST",
            route: "/api/user/{id}/follow",
            response: "No content",
        },
        {
            method: "POST",
            route: "/api/user/{id}/followers",
            response: [
                {
                    id: "string",
                    nickname: "string",
                    email: "string",
                    firstName: "string",
                    lastName: "string",
                    dateOfBirth: "string",
                    imagePath: "string | null",
                    aboutMe: "string | null",
                    private: "bool",
                    timestamp: "string",
                },
            ],
        },
    ];

    return (
        <div id="container">
            <aside></aside>
            <main>
                <ul>
                    {routes.map((route, index) => (
                        <li key={index}>
                            <span className="method">{route.method}</span>
                            <code>{route.route}</code>
                            {route.body && (
                                <div className="request">
                                    <div>Request:</div>
                                    <pre>
                                        {JSON.stringify(route.body, null, 2)}
                                    </pre>
                                </div>
                            )}
                            <div className="response">Response:</div>
                            <pre>{JSON.stringify(route.response, null, 2)}</pre>
                        </li>
                    ))}
                </ul>
            </main>
        </div>
    );
}
