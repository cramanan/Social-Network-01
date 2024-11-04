import { ClientChat } from "@/types/chat";
import { createContext, useContext } from "react";

type WebSocketContextType = {
    socket: WebSocket;
    sendChat: (chat: ClientChat) => void;
};

export const webSocketContext = createContext<WebSocketContextType | null>(
    null
);
