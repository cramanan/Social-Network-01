export type Post = {
    id: string;
    userId: string;
    username: string;
    userImage: string;
    groupId: string | null;
    content: string;
    images: string[];
    privacyLevel: "public" | "private" | "almost_private";
    selectedUserIds: string[];
    timestamp: string;
};

export type Comment = {
    userImage: string;
    username: string;
    userId: string;
    postId: string;
    content: string;
    image: string;
    timestamp: string;
};
