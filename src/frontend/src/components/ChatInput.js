import React, { useState } from "react";
import "../styles/ChatInput.css";

const ChatInput = ({ addMessage }) => {
  const [message, setMessage] = useState("");

  const handleMessageChange = (event) => {
    setMessage(event.target.value);
  };

  const handleSendMessage = () => {
    if (message.trim() !== "") {
      addMessage({ content: message, fromMe: true });
      setMessage("");
    }
  };

  return (
    <div className="chat-input">
      <input
        type="text"
        placeholder="Send a message..."
        value={message}
        onChange={handleMessageChange}
      />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
};

export default ChatInput;
