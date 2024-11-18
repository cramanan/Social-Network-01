export type SocketMessageType = "message" | "ping" | "friend-request";

export type SocketMessage<T = unknown> = {
    type: SocketMessageType;
    data: T;
};

export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & { timestamp: string };
