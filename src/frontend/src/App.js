import './App.css';
import './styles/LandingPage.css';
import React, { useState } from 'react';
import ChatList from "./components/ChatList";
import ChatInput from "./components/ChatInput";
import ChatItem from "./components/ChatItem";
import ChatHistory from './components/ChatHistory';
import RadioButton from './components/RadioButton';

function App() {
  // const [messages, setMessages] = useState({ user: [], bot: [] });
  const [messages, setMessages] = useState([]);
  const [message, setMessage] = useState("");

const addMessage = (message) => {
  setMessages((messages) => [
    ...messages,
    { ...message, fromMe: true }
  ]);
  
  setTimeout(() => {
    // Remove content from message
    setMessages((messages) => [
      ...messages,
      {
        botMessage: "Ada yang bisa saya bantu?",
        botName: "Bot",
        isBot: true,
        fromMe: false
      }
    ]);
    setMessage("");
  }, 1000);
};

  const clearScreen = () => {
    setMessages([]);
  };


  return (
    <div className="myBG">
      <div className="container">
        <div className="chat-history"></div>
        <h1 className="title">ChatterBot-STIM(AI)</h1>
        <h2 className="history">History</h2>
        <h3 className="algorithms">Algoritma:</h3>
        <ChatList messages={messages} />
        <ChatInput addMessage={addMessage} />
        <ChatHistory chatHistory={clearScreen} />
        <RadioButton
          name="algoritma"
          value1="0"
          label1="   KMP"
          value2="1"
          label2="   BM"
          defaultValue="0"
        />
      </div>
    </div>
  );
}

export default App;
