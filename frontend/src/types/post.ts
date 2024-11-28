export type Post = {
    id: string;
    username: string;
    userId: string;
    groupId: string | null;
    content: string;
    images: string[];
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
