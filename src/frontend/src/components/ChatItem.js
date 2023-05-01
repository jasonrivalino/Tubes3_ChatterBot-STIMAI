import React from "react";
import "../styles/ChatItem.css";
import zull from "./zul.jpg";
import robot from "./robot.jpeg";
import "./ChatInput.js"

const ChatItem = ({ message }) => {
  return (
    <div className={message.fromMe ? "from-me" : "from-them"}>
      <img src={zull} className="chat-avatar" alt="zull" />
      <div className="chat-bubble">
        <div className="chat-username">{message.fromMe ? "Jason" : message.username}</div><br/>
        <div className="chat-content">{message.content}</div>
      </div>
      {message.isBot && (
        <div className="chat-bubble-bot">
          <img src={robot} className="chat-avatar" alt="zull" />
          <div className="chat-bubble-bot">
            <div className="chat-username">{message.botName}</div><br/>
            <div className="chat-content">{message.botMessage}</div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ChatItem;
