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
<<<<<<< HEAD
=======

>>>>>>> 5d6472f9912a0792d5438214c53a2b21a5cd2de5
        <div className="chat-content">{message.content}</div>
      </div>
      {message.isBot && (
        <div className="chat-bubble-bott">
          <img src={robot} className="chat-avatar-bot" alt="zull" />
          <div className="chat-bubble-bot">
            <div className="chat-username-bot">{message.botName}</div><br/>
            <div className="chat-content-bot">{message.botMessage}</div>
            </div>         
        </div>
      )}
    </div>
  );
};

export default ChatItem;