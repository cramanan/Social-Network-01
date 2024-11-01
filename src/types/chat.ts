export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & {
    timestamp: string;
};
