export type SocketMessageType = "message" | "ping" | "follow-request";

export type SocketMessage<T = unknown> = {
    type: SocketMessageType;
    data: T;
};

export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & { senderid: string; timestamp: string };
