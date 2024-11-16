export type SocketMessage<T = unknown> = {
    type: "message" | "ping" | "friend-request";
    data: T;
};

export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & { timestamp: string };
