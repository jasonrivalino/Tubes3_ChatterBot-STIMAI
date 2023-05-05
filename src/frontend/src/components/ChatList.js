
import React, { useRef, useEffect } from "react";
import "../styles/ChatList.css";
import ChatItem from "./ChatItem";

const ChatList = ({ messages }) => {
  const chatEnd = useRef(null);

  useEffect(() => {
    chatEnd.current.scrollIntoView({ behavior: "smooth" });
  });

  return (
    <div className="chat-list">
      {messages.map((message, index) => (
        <ChatItem key={index} message={message} />
      ))}
      <div ref={chatEnd}></div>
    </div>
  );
};

export default ChatList;

