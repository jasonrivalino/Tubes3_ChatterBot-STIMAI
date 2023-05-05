import "../styles/ChatHistory.css";
import { useState } from "react";

const ChatHistory = ({chatHistory}) => {

  const [chatCount, setChatCount] = useState(0);

  const handleAddChat = () => {
    chatHistory();
    setChatCount((count) => count + 1);
  };

  return (
    <>
      <div className="chat-history-button">
        <button onClick={handleAddChat}>+ Add New Chat</button>
      </div>
      <div className="historyButtons-container">
        {[...Array(chatCount)].map((_, index) => (
          <div key={index} className="historyButton-wrapper">
            <button className="historyButtons">History {index + 1}</button>
          </div>
        ))}
        <div ref={(el) => {el?.scrollIntoView({ behavior: "smooth" });}}></div>
      </div>
    </>
  );
};

export default ChatHistory;