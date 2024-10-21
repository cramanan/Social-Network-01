import { createContext, useContext } from "react";

export const webSocketContext = createContext<WebSocket | null>(null);

export const useWebSocket = () => useContext(webSocketContext);
