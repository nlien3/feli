// src/components/MyCourses.js
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import axios from "axios";
import {
  Container,
  Grid,
  Card,
  CardContent,
  CardMedia,
  Typography,
  CircularProgress,
  Alert,
  Box,
  Button,

} from "@mui/material";

const MyCourses = () => {
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(() => {
    const fetchCourses = async () => {
      const userId = localStorage.getItem("userId"); // Obtener el userId desde localStorage
      if (!userId) {
        setError(true);
        setLoading(false);
        return;
      }

      try {
        const response = await axios.get(
          `${process.env.REACT_APP_API}/subscriptions/${userId}`,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`,
              "Content-Type": "application/json",
            },
          }
        ); // Reemplaza con tu URL de la API
        setCourses(response.data);
      } catch (error) {
        console.error("Error fetching courses:", error);
        setError(true);
      } finally {
        setLoading(false);
      }
    };

    fetchCourses();
  }, []);

  if (loading) {
    return <CircularProgress />;
  }

  if (error) {
    return (
      <Container>
        <Alert severity="error">
          Ha ocurrido un error al cargar los cursos.
        </Alert>
      </Container>
    );
  }

  if (courses.length === 0) {
    return (
      <Container>
        <Typography variant="h6" align="center" gutterBottom>
          Aún no te has inscripto a ningún curso.
        </Typography>
      </Container>
    );
  }

  return (
    <Container>
      <Typography variant="h4" gutterBottom>
        Mis Cursos:
      </Typography>
      <Grid container spacing={4}>
        {courses.map((subscription) => (
          <Grid item key={subscription.IdSubscription} xs={12} sm={6} md={4}>
            <Card>
              <CardMedia
                component="img"
                height="140"
                image={subscription.curso.imageURL}
                alt={subscription.curso.nombre}
              />
              <CardContent>
                <Typography variant="h5" component="div">
                  {subscription.curso.nombre}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {subscription.curso.descripcion}
                </Typography>
                <Typography variant="body1" color="text.primary">
                  Precio: ${subscription.curso.precio}
                </Typography>
                <Typography variant="body1" color="text.primary">
                  Categoría: {subscription.curso.categoria}
                </Typography>
                <Typography variant="body1" color="text.primary">
                  Dificultad: {subscription.curso.dificultad}
                </Typography>
                <Box mt={2}>
                  <Button
                    component={Link}
                    to={`/my-courses/${subscription.curso.ID}`}
                    variant="contained"
                    color="primary"
                  >
                    Más Información
                  </Button>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default MyCourses;
