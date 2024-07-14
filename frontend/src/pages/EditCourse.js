import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
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

const EditCourse = () => {
  const { id } = useParams();
  const navigate = useNavigate();
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

  useEffect(() => {
    const fetchCourseData = async () => {
      try {
        const response = await axios.get(
          `${process.env.REACT_APP_API}/courses/${id}`
        );
        setCourseData({
          Nombre: response.data.nombre,
          Categoria: response.data.categoria,
          Dificultad: response.data.dificultad,
          Precio: response.data.precio,
          Descripcion: response.data.descripcion,
          ImageURL: response.data.imageURL,
        });
      } catch (error) {
        console.error("Error fetching course data:", error);
        setError("Error fetching course data.");
      }
    };

    fetchCourseData();
  }, [id]);

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
      await axios.put(
        `${process.env.REACT_APP_API}/courses/${id}`,
        courseData,
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        }
      );
      setSuccess(true);
      setError("");
      setTimeout(() => {
        navigate("/"); 
      }, 2000);
    } catch (error) {
      console.error("Error updating course:", error);
      setError("Ha ocurrido un error al editar el curso.");
      setSuccess(false);
    }
  };

  return (
    <Container maxWidth="sm">
      <Box mt={5}>
        <Typography variant="h4" component="h1" gutterBottom>
          Editar Curso
        </Typography>
        {error && <Alert severity="error">{error}</Alert>}
        {success && (
          <Alert severity="success">
            El curso ha sido editado exitosamente.
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
              Editar Curso
            </Button>
          </Box>
        </form>
      </Box>
    </Container>
  );
};

export default EditCourse;
