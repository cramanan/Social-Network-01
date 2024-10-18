import type { NextRequest } from "next/server";

export async function middleware(request: NextRequest) {
    console.log("Middleware triggered for:", request.nextUrl.pathname);
    try {
        const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_URL}/api/auth`,
            {
                headers: request.headers,
            }
        );
        console.log("API response status:", response.status);

        if (!response.ok)
            return Response.redirect(new URL("/auth", request.url));
    } catch (error) {
        console.error(error);
        return Response.redirect(new URL("/auth", request.url));
    }
}

export const config = {
    matcher: ["/((?!api|_next/static|_next/image|auth))"],
};
