export type Post = {
    id: string;
    userId: string;
    username: string;
    userImage: string;
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
