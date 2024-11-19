import { ClientChat } from "@/types/chat";
import { createContext } from "react";

type WebSocketContextType = {
    socket: WebSocket;
    sendChat: (chat: ClientChat) => void;
};

export const webSocketContext = createContext<WebSocketContextType>({
    socket: new WebSocket(`ws://${process.env.NEXT_PUBLIC_API_URL}/api/socket`),
    sendChat: () => {},
});
