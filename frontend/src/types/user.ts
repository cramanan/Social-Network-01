export type User = {
    id: string;
    email: string;
    nickname: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
    image: string;
    aboutMe?: string;
    private: boolean;
};

export type OnlineUser = User & { online: boolean };
