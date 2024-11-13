import type { NextRequest } from "next/server";

export async function middleware(request: NextRequest) {
    try {
        const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_URL}/api/profile`,
            {
                headers: request.headers,
            }
        );
        if (!response.ok)
            return Response.redirect(new URL("/auth", request.url));
    } catch (error) {
        console.error(error);
        return Response.redirect(new URL("/auth", request.url));
    }
}
export const config = {
    matcher: ["/((?!api|_next/static|_next/image|.*\\.png$|auth).*)"],
};
