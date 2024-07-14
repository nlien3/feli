import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  Box,
  Button,
  Container,
  TextField,
  Typography,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";

const Chat = ({ courseId }) => {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchMessages();
  }, []);

  const fetchMessages = async () => {
    try {
      const response = await axios.get(
        `${process.env.REACT_APP_API}/course/${courseId}/chat`,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`,
              "Content-Type": "application/json",
            },
          }
      );
      setMessages(response.data);
      setLoading(false);
    } catch (error) {
      console.error("Error fetching chat messages:", error);
      setLoading(false);
    }
  };

  const handleSendMessage = async () => {
    if (newMessage.trim() === "") return;

    try {
      const userId = localStorage.getItem("userId"); // Assuming the user ID is stored in localStorage
      const response = await axios.post(
        `${process.env.REACT_APP_API}/course/chat`,
        {
          IdUsuario: parseInt(userId),
          IdCurso: courseId,
          Message: newMessage,
        },
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`,
              "Content-Type": "application/json",
            },
          }
      );

      setMessages([response.data, ...messages]);
      setNewMessage("");
    } catch (error) {
      console.error("Error sending message:", error);
    }
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" gutterBottom>
        Chat
      </Typography>
      <Box
        component="form"
        onSubmit={(e) => {
          e.preventDefault();
          handleSendMessage();
        }}
      >
        <TextField
          label="New Message"
          variant="outlined"
          fullWidth
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
        />
        <Button
          variant="contained"
          color="primary"
          fullWidth
          onClick={handleSendMessage}
        >
          Send
        </Button>
      </Box>
      <List>
        {loading ? (
          <Typography>Loading...</Typography>
        ) : (
          messages.map((message) => (
            <ListItem key={message.IdChat}>
              <ListItemText
                primary={message.NombreUsuario}
                secondary={message.Message}
              />
            </ListItem>
          ))
        )}
      </List>
    </Container>
  );
};

export default Chat;
