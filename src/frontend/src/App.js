import './App.css';
import './styles/LandingPage.css';
import React, { useState } from "react";
import ChatList from "./components/ChatList";
import ChatInput from "./components/ChatInput";
import "./styles/ChatHistory.css";
import RadioButton from './components/RadioButton';

function App() {
  const [messages, setMessages] = useState([]);
  const addMessage = (message) => {
    setMessages([...messages, {...message, fromMe: true}]);
    setTimeout(() => {
      if (message.fromMe) {
        setMessages([
          ...messages,
          {
            botMessage: "Ada yang bisa saya bantu?",
            botName: "Bot",
            isBot: true,
            fromMe: false,
          },
        ]);
      }
    }, 1500);
  };
  

  return (
    <div className="myBG">
    <div className="container">
      <div className='chat-history'></div>
      <h1 className="title">ChatterBot-STIM(AI)</h1>
      <h2 className="history">History</h2>
      <h3 className="algorithms">Algoritma:</h3>
      <ChatList messages={messages} />
      <ChatInput addMessage={addMessage} />
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
