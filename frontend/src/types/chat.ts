export type SocketMessage<T = unknown> = {
    type : "message" | "ping"
    data : T
}

export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & {
    timestamp: string;
};
