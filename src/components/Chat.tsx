"use client";

import React, { useState } from "react";
import { UserListIcon } from "./icons/UserListIcon";
import UserList from "./UserList";
import FindUser from "./FindUser";
import ChatList from "./ChatList";
import { ChatIcon } from "./icons/ChatIcon";
import { FindUserIcon } from "./icons/FindUserIcon";

const Chat = () => {
  const useToggle = (initialState = false) => {
    const [state, setState] = useState(initialState);
    const toggle = () => setState((prev) => !prev);
    return [state, toggle] as const;
  };

  const [isChatListOpen, toggleChatList] = useToggle(false);
  const [isFindUserOpen, toggleFindUser] = useToggle(false);
  const [isUserListOpen, toggleUserList] = useToggle(false);

  const handleChatIconClick = () => {
    toggleChatList()
    if (isFindUserOpen) toggleFindUser();
    if (isUserListOpen) toggleUserList();
  };

  const handleFindUserIconClick = () => {
    toggleFindUser();
    if (isChatListOpen) toggleChatList();
    if (isUserListOpen) toggleUserList();
  };

  const handleUserListIconClick = () => {
    toggleUserList();
    if (isChatListOpen) toggleChatList();
    if (isFindUserOpen) toggleFindUser();
  };

  const navItems = [
    {
      icon: <ChatIcon />,
      OnClick: handleChatIconClick,
      label: "Chat list",
      expanded: isChatListOpen,
    },
    {
      icon: <FindUserIcon />,
      OnClick: handleFindUserIconClick,
      label: "FindUser list",
      expanded: isFindUserOpen,
    },
    {
      icon: <UserListIcon />,
      OnClick: handleUserListIconClick,
      label: "User list",
      expanded: isUserListOpen,
    },
  ];

  return (
    <>
      <nav
        id="chat-nav"
        className="w-72 flex bg-white bg-opacity-40 m-3 rounded-b-3xl px-7"
        aria-label="Chat Navigation"
      >
        <ul className="w-full flex flex-row justify-between items-center">
          {navItems.map((items, index) => (
            <li key={index}>
              <button
                onClick={items.OnClick}
                aria-label={`Toggle ${items.label}`}
                aria-expanded={items.expanded}
              >
                <span className="sr-only">{items.label}</span>
                {items.icon}
              </button>
            </li>
          ))}
        </ul>
      </nav>

      <div
        id="chat-list"
        aria-hidden={!isChatListOpen}
        className={`
        absolute top-16 right-[3.4rem] transition-all duration-300 ease-in-out
        ${isChatListOpen
            ? "opacity-100 translate-y-0 pointer-events-auto"
            : "opacity-0 translate-y-5 pointer-events-none"
          }
      `}
        aria-label="Chat list"
      >
        <ChatList />
      </div>

      <div
        id="finduser-list"
        aria-hidden={!isFindUserOpen}
        className={`
        absolute right-3 top-16 transition-all duration-300 ease-in-out
        ${isFindUserOpen
            ? "opacity-100 translate-y-0 pointer-events-auto"
            : "opacity-0 translate-y-5 pointer-events-none"
          }
      `}
        aria-label="Find user list"
      >
        <FindUser />
      </div>

      <div
        id="user-list"
        aria-hidden={!isUserListOpen}
        className={`
        absolute right-3 top-16 transition-all duration-300 ease-in-out
        ${isUserListOpen
            ? "opacity-100 translate-y-0 pointer-events-auto"
            : "opacity-0 translate-y-5 pointer-events-none"
          }
      `}
        aria-label="User list"
      >
        <UserList />
      </div>
    </>
  );
};

export default Chat;
