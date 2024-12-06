export type SocketMessageType =
    | "message"
    | "ping"
    | "follow-request"
    | "group-request"
    | "group-invite"
    | "event";

export type SocketMessage<T = unknown> = {
    type: SocketMessageType;
    data: T;
};

export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & { senderId: string; timestamp: string };
