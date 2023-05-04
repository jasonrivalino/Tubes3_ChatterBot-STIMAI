import "../styles/ChatHistory.css";

const ChatHistory = ({ chatHistory }) => {
    const clearScreen = () => {
        chatHistory();
    };

    return(
        <div className="chat-history-button">
            <button onClick={clearScreen}>+ㅤAdd New Chat</button>
        </div>
        
    );
};

export default ChatHistory;