"use client";

import React, { useEffect, useState } from "react";
import { BackIcon } from "./icons/BackIcon";
import { SendIcon } from "./icons/SendIcon";
import { EmoteIcon } from "./icons/EmoteIcon";
import EmojiPicker, { EmojiClickData } from "emoji-picker-react";
import { User } from "@/types/user";
import { ClientChat, ServerChat, SocketMessage } from "@/types/chat";
import { useWebSocket } from "@/hooks/useWebSocket";
import Image from "next/image";

interface ChatBoxProps {
    onClose?: () => void;
    recipient: User;
}

const ChatBox = ({ onClose, recipient }: ChatBoxProps) => {
    const [messages, setMessages] = useState<ServerChat[]>([]);
    const [EmojiPick, setEmojiPicker] = useState(false);
    const [chat, setChat] = useState<ClientChat>({
        recipientId: recipient.id,
        content: "",
    });

    useEffect(() => {
        const fetchMessages = async () => {
            const response = await fetch(`/api/users/${recipient.id}/chats`);
            const data = await response.json();
            setMessages(data);
        };

        fetchMessages();
    }, [recipient.id]);

    const websocket = useWebSocket();

    useEffect(() => {
        const addMessage = (msg: MessageEvent) => {
            const message = JSON.parse(msg.data) as SocketMessage<ServerChat>;
            if (message.type !== "message") return;
            setMessages((prev) => [...prev, message.data]);
        };

        websocket.socket.addEventListener("message", addMessage);
        return () =>
            websocket.socket.removeEventListener("message", addMessage);
    }, [websocket]);

    const HandleEmoji = (emojiData: EmojiClickData) => {
        setChat((prev) => ({
            ...prev,
            content: prev.content + emojiData.emoji,
        }));
        setEmojiPicker(false);
    };

    if (!websocket) return <>No socket</>;

    return (
        <div className="flex flex-col w-full h-full relative xl:w-[343px] xl:rounded-[25px] xl:h-[642px] xl:bg-[#fbfbfb]">
            <div className="flex flex-row w-full min-h-14 items-center justify-between border-b border-black px-3 xl:rounded-tl-[25px] xl:rounded-t-[25px] xl:bg-[#445ab3]/20 xl:w-[343px]">
                <button onClick={onClose}>
                    <BackIcon />
                </button>
                <span>{recipient.nickname}</span>
                <Image
                    src={recipient.image}
                    alt=""
                    width={40}
                    height={40}
                    className="w-9 h-9 border border-black rounded-full"
                />
            </div>

            <ul className="flex flex-col flex-grow px-3 py-2 overflow-scroll no-scrollbar">
                {messages.map((msg, index) => {
                    const isRecipient = msg.recipientId === recipient.id;
                    const timestamp = new Date(msg.timestamp);

                    return (
                        <li
                            key={index}
                            className={`flex flex-col ${
                                isRecipient
                                    ? "self-end items-end"
                                    : "self-start"
                            }`}
                        >
                            <p
                                className={`p-3 rounded-2xl w-fit ${
                                    isRecipient
                                        ? "bg-[#b88ee5] text-black"
                                        : "bg-[#4174e2] text-white"
                                }`}
                            >
                                {msg.content}
                            </p>
                            <div>
                                {timestamp.toLocaleDateString()},
                                {timestamp.toLocaleTimeString()}
                            </div>
                        </li>
                    );
                })}
            </ul>

            {EmojiPick && (
                <div className="absolute bottom-20 left-0">
                    <EmojiPicker onEmojiClick={HandleEmoji} />
                </div>
            )}

            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    websocket.sendChat(chat);
                    setMessages((prev) => [
                        ...prev,
                        {
                            senderId: "",
                            timestamp: new Date().toString(),
                            ...chat,
                        },
                    ]);
                    setChat({ ...chat, content: "" });
                }}
                className="h-[50px] flex flex-row items-center m-5 bg-[#445ab3]/20 rounded-[25px] p-2 gap-2"
            >
                <button
                    type="button"
                    onClick={() => setEmojiPicker(!EmojiPick)}
                >
                    <EmoteIcon />
                </button>
                <input
                    type="text"
                    placeholder="Enter your message"
                    onChange={(e) =>
                        setChat({ ...chat, content: e.target.value })
                    }
                    value={chat.content}
                    className="bg-white/0 w-full placeholder:text-black"
                />
                <button type="submit">
                    <SendIcon />
                </button>
            </form>
        </div>
    );
};

export default ChatBox;
