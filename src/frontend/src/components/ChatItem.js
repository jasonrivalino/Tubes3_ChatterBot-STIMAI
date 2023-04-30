import React from "react";
import "../styles/ChatItem.css";

const ChatItem = ({ message }) => {
  return (
    <div className={`chat-item ${message.fromMe ? "from-me" : ""}`}>
      <div className="chat-content">{message.content}</div>
    </div>
  );
};

export default ChatItem;
