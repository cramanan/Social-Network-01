import { StrictOmit } from "@/utils/types";

export type User = {
    id: string;
    email: string;
    nickname: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
    image: string;
    aboutMe: string | null;
    isPrivate: boolean;
};

export type OnlineUser = User & { online: boolean };

export type EditableUser = StrictOmit<User, "id" | "dateOfBirth" | "email">;
