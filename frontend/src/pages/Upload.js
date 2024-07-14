import React, { useState } from "react";
import axios from "axios";
import { Box, Button, Container, TextField, Typography } from "@mui/material";

const Upload = () => {
  const [file, setFile] = useState(null);
  const [message, setMessage] = useState("");

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file) {
      setMessage("Please select a file first");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API}/upload`,
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      setMessage(response.data.message);
    } catch (error) {
      setMessage("Failed to upload file");
      console.error("Error uploading file:", error);
    }
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" gutterBottom>
        Upload File
      </Typography>
      <Box>
        <TextField type="file" fullWidth onChange={handleFileChange} />
        <Button
          variant="contained"
          color="primary"
          fullWidth
          onClick={handleUpload}
        >
          Upload
        </Button>
      </Box>
      {message && <Typography variant="body1">{message}</Typography>}
    </Container>
  );
};

export default Upload;
