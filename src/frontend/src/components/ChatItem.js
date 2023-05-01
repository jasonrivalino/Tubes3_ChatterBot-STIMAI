import React from "react";
import "../styles/ChatItem.css";
import zull from "./zul.jpg";

const ChatItem = ({ message }) => {
  return (
    // Set message from me and get return from program
    <div className= {message.fromMe ? "from-me" : "from-them"}>
      <div className="chat-bubble">
      <img src = {zull} className="chat-avatar" alt="zull" />
        <div className="chat-username">
          {message.fromMe ? "Zulfiansyah" : message.username}
        </div><br/>
      </div>
      <div className="chat-content">{message.content}</div>
    </div>
  );
};

export default ChatItem;
