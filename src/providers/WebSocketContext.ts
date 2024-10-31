import { Chat } from "@/types/chat";
import { createContext, useContext } from "react";

type WebSocketContextType = {
    socket: WebSocket;
    sendChat: (chat: Chat) => void;
    ping: () => void;
};

export const webSocketContext = createContext<WebSocketContextType | null>(
    null
);

export const useWebSocket = () => useContext(webSocketContext);
