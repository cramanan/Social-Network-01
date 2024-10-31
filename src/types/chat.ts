export type Chat = {
    recipientId: string;
    content: string;
};

export type ServerChat = Chat & {
    timestamp: string;
};
