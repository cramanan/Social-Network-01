import { webSocketContext } from "@/providers/WebSocketContext";
import { useContext } from "react";

export const useWebSocket = () => useContext(webSocketContext);
