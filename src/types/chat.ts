type SocketMessage<T = unknown> = {
    type : string
    data : T
}

export type ClientChat = {
    recipientId: string;
    content: string;
};

export type ServerChat = ClientChat & {
    timestamp: string;
};
