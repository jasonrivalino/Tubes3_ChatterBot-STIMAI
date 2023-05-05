import React from "react";
import "../styles/ChatItem.css";
import profil from "./assets/blank.jpg";
import robot from "./assets/robot.jpeg";
import "./ChatInput.js"

const ChatItem = ({ message }) => {
  return (
    <div className={message.fromMe ? "from-me" : "from-them"}>
      {message.fromMe && (
        <img src={profil} className="chat-avatar" alt="profil" />
      )}
      {!message.fromMe && message.isBot ? (
        <div className="chat-bubble-bott">
          <img src={robot} className="chat-avatar-bot" alt="robot" />
          <div className="chat-bubble-bot">
            <div className="chat-username-bot">{message.botName}</div>
            <br />
            <div className="chat-content-bot" style={{ wordBreak: "break-word" }}>{message.botMessage}</div>
          </div>
        </div>
      ) : (
        <div className="chat-bubble">
          <div className="chat-username">{message.fromMe ? "Jason" : message.username}</div>
          <br />
          <div className="chat-content" style={{ wordBreak: "break-word" }}>{message.content}</div>
        </div>
      )}
    </div>
  );
};

export default ChatItem;
