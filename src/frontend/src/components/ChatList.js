import React from "react";
import "../styles/ChatList.css";
import ChatItem from "./ChatItem";

const ChatList = ({ messages }) => {
  return (
    <div className="chat-list">
      {messages.map((message, index) => (
        <ChatItem key={index} message={message} />
      ))}
      <div ref={(el) => {el?.scrollIntoView({ behavior: "smooth" });}}></div>
    </div>
  );
};

export default ChatList;
