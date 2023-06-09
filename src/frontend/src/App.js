import './App.css';
import './styles/LandingPage.css';
import React, { useState } from 'react';
import ChatList from "./components/ChatList";
import ChatInput from "./components/ChatInput";
import ChatHistory from './components/ChatHistory';
import RadioButton from './components/RadioButton';
import axios from 'axios';

function App() {
  const [messages, setMessages] = useState([]);
  const [message, setMessage] = useState("");
  const [selectedAlgorithm, setSelectedAlgorithm] = useState("1"); // Menyimpan nilai dari RadioButton

  const addMessage = (message) => {
    setMessages((messages) => [
      ...messages,
      { ...message, fromMe: true }
    ]);

    // SetTimeOut dengan botMessage berasal dari response Post API
    setTimeout(() => {
      axios.post('http://localhost:8080/api/history', {
        id_riwayat : 3,
        pertanyaan : message.content,
        str_match : parseInt(selectedAlgorithm) // Menggunakan nilai dari RadioButton
      })
      .then(function (response) {
        console.log(response);
        setMessages((messages) => [
          ...messages,
          {
            botMessage: response.data.jawaban,
            botName: "Bot",
            isBot: true,
            fromMe: false
          }
        ]);
        setMessage("");
      })
      .catch(function (error) {
        console.log(error);
      });
    }, 1000);
  };

  const clearScreen = () => {
    setMessages([]);
  };

  const handleAlgorithmChange = (value) => {
    setSelectedAlgorithm(value);
  };
  return (
    <div className="myBG">
      <div className="container">
        <div className="chat-history"></div>
        <h1 className="title">ChatterBot-STIM(AI)</h1>
        <h1 className="nameKelompok">Jason (13521008), Syauqi (13521014), Zul (13521028)</h1>
        <h2 className="history">History</h2>
        <div className="radioButtonMenu"></div>
        <h3 className="algorithms">Algoritma:</h3>
        <ChatList messages={messages} />
        <ChatInput addMessage={addMessage} />
        <ChatHistory chatHistory={clearScreen} />
        <RadioButton
          name="algoritma"
          value1="1"
          label1="   KMP"
          value2="2"
          label2="   BM"
          defaultValue="1"
          onChange={handleAlgorithmChange}
        />
      </div>
    </div>
  );
}

export default App;
