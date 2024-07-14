import React, { useState } from "react";
import axios from "axios";
import {
  Container,
  TextField,
  Button,
  Typography,
  Alert,
  Box,
  MenuItem,
  Select,
  FormControl,
  InputLabel,
} from "@mui/material";

const CreateCourse = () => {
  const [courseData, setCourseData] = useState({
    Nombre: "",
    Categoria: "",
    Dificultad: "",
    Precio: "",
    Descripcion: "",
    ImageURL: "",
  });
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCourseData({
      ...courseData,
      [name]: name === "Precio" ? parseFloat(value) : value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const token = localStorage.getItem("token");
    if (!token) {
      setError("No se encontró un token de autenticación.");
      return;
    }

    try {
      await axios.post(`${process.env.REACT_APP_API}/courses`, courseData, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });
      setSuccess(true);
      setError("");
      setCourseData({
        Nombre: "",
        Categoria: "",
        Dificultad: "",
        Precio: "",
        Descripcion: "",
        ImageURL: "",
      });
    } catch (error) {
      console.error("Error creating course:", error);
      setError("Ha ocurrido un error al crear el curso.");
      setSuccess(false);
    }
  };

  return (
    <Container maxWidth="sm">
      <Box mt={5}>
        <Typography variant="h4" component="h1" gutterBottom>
          Crear Nuevo Curso
        </Typography>
        {error && <Alert severity="error">{error}</Alert>}
        {success && (
          <Alert severity="success">
            El curso ha sido creado exitosamente.
          </Alert>
        )}
        <form onSubmit={handleSubmit}>
          <TextField
            label="Nombre"
            name="Nombre"
            variant="outlined"
            fullWidth
            margin="normal"
            value={courseData.Nombre}
            onChange={handleChange}
            required
          />
          <TextField
            label="Categoría"
            name="Categoria"
            variant="outlined"
            fullWidth
            margin="normal"
            value={courseData.Categoria}
            onChange={handleChange}
            required
          />
          <FormControl fullWidth margin="normal" variant="outlined" required>
            <InputLabel>Dificultad</InputLabel>
            <Select
              label="Dificultad"
              name="Dificultad"
              value={courseData.Dificultad}
              onChange={handleChange}
              required
            >
              <MenuItem value="Facil">Facil</MenuItem>
              <MenuItem value="Medio">Medio</MenuItem>
              <MenuItem value="Dificil">Dificil</MenuItem>
            </Select>
          </FormControl>
          <TextField
            label="Precio"
            name="Precio"
            type="number"
            variant="outlined"
            fullWidth
            margin="normal"
            value={courseData.Precio}
            onChange={handleChange}
            required
          />
          <TextField
            label="Descripción"
            name="Descripcion"
            variant="outlined"
            fullWidth
            margin="normal"
            multiline
            rows={4}
            value={courseData.Descripcion}
            onChange={handleChange}
            required
          />
          <TextField
            label="URL de la Imagen"
            name="ImageURL"
            variant="outlined"
            fullWidth
            margin="normal"
            value={courseData.ImageURL}
            onChange={handleChange}
            required
          />
          <Box mt={2}>
            <Button type="submit" variant="contained" color="primary" fullWidth>
              Crear Curso
            </Button>
          </Box>
        </form>
      </Box>
    </Container>
  );
};

export default CreateCourse;
