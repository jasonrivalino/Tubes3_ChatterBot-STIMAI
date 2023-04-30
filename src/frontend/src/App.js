import './App.css';
import './styles/LandingPage.css';
import React, { useState } from "react";
import ChatList from "./components/ChatList";
import ChatInput from "./components/ChatInput";

function App() {
  const [messages, setMessages] = useState([]);
  const addMessage = (message) => {
    setMessages([...messages, message]);
  };

  return (
    <div className="myBG">
    <div className="container">
      <ChatList messages={messages} />
      <ChatInput addMessage={addMessage} />
    </div>
    </div>
  );
}

export default App;
