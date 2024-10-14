import type { NextRequest } from "next/server";

export async function middleware(request: NextRequest) {
    try {
        const response = await fetch(`${process.env.API_URL}/api/auth`, {
            headers: request.headers,
        });

        if (!response.ok)
            return Response.redirect(new URL("/auth", request.url));
    } catch (error) {
        return Response.redirect(new URL("/auth", request.url));
    }
}

export const config = {
    matcher: ["/((?!api|_next/static|_next/image|.*\\.png$|auth).*)"],
};
