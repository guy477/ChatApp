import React, { useState } from 'react';

function App() {
  // Basic user inputs for demonstration
  const [userId, setUserId] = useState<string>("user_123");
  const [chatId, setChatId] = useState<string>("");
  const [userMessage, setUserMessage] = useState<string>("");
  const [conversation, setConversation] = useState<string>("");

  const handleSend = async () => {
    try {
      const payload = {
        user_id: userId,
        chat_id: chatId,
        message: userMessage
      };
      
      // Add the user message immediately to the conversation
      setConversation(prev => prev + `User: ${userMessage}\n`);
      
      // Create EventSource for SSE
      const response = await fetch("http://localhost:3000/api/chat", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
      });

      // Create a new line for assistant's response
      setConversation(prev => prev + "Assistant: ");

      const reader = response.body?.getReader();
      if (!reader) {
        throw new Error("ReadableStream not supported");
      }

      const decoder = new TextDecoder();
      while (true) {
        const {value, done} = await reader.read();
        if (done) break;
        
        const text = decoder.decode(value);
        // Split the text into SSE events
        const events = text.split('\n\n').filter(Boolean);
        
        for (const event of events) {
          if (event.startsWith('data: ')) {
            const content = event.slice(6); // Remove 'data: ' prefix
            setConversation(prev => prev + content);
          } else if (event.startsWith('event: error')) {
            const errorMsg = event.split('\n')[1]?.slice(6); // Get error message
            throw new Error(errorMsg || 'Unknown error');
          }
        }
      }

      // Add a line break after assistant's complete response
      setConversation(prev => prev + "\n\n");
      setUserMessage(""); // Clear input field

    } catch (error: any) {
      console.error("Error in handleSend:", error);
      alert(error.message || error.toString());
    }
  };

  return (
    <div style={{ maxWidth: 600, margin: "auto", padding: 16 }}>
      <h1>Mini Chat Demo (TS + React + Go Service)</h1>

      <div style={{ marginBottom: 8 }}>
        <label>User ID:</label>
        <input 
          type="text"
          value={userId}
          onChange={(e) => setUserId(e.target.value)}
          style={{ width: "100%", marginTop: 4 }}
        />
      </div>

      <div style={{ marginBottom: 8 }}>
        <label>Chat ID (optional):</label>
        <input 
          type="text"
          value={chatId}
          onChange={(e) => setChatId(e.target.value)}
          placeholder="Leave blank to create a new chat"
          style={{ width: "100%", marginTop: 4 }}
        />
      </div>

      <div style={{ marginBottom: 8 }}>
        <label>Message:</label>
        <textarea 
          value={userMessage}
          onChange={(e) => setUserMessage(e.target.value)}
          style={{ width: "100%", height: 60, marginTop: 4 }}
        />
      </div>

      <button onClick={handleSend} style={{ marginBottom: 16 }}>
        Send to Go
      </button>

      <h2>Conversation</h2>
      <div style={{ 
          border: "1px solid #ccc", 
          padding: 8, 
          height: 300, 
          overflowY: "auto", 
          whiteSpace: "pre-wrap" 
        }}>
        {conversation}
      </div>
    </div>
  );
}

export default App;