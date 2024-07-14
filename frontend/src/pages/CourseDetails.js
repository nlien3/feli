import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import {
  Container,
  Card,
  CardContent,
  CardMedia,
  Typography,
  CircularProgress,
  Alert,
} from "@mui/material";
import Chat from "./Chat";
import Upload from "./Upload";

const CourseDetails = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await axios.get(
          `${process.env.REACT_APP_API}/courses/suscription/${id}`,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`,
              "Content-Type": "application/json",
            },
          }
        ); // Reemplaza con tu URL de la API
        setCourse(response.data);
      } catch (error) {
        console.error("Error fetching course:", error);
        setError(true);
      } finally {
        setLoading(false);
      }
    };

    fetchCourse();
  }, [id]);

  if (loading) {
    return <CircularProgress />;
  }

  if (error) {
    return (
      <Container>
        <Alert severity="error">Ha ocurrido un error al cargar el curso.</Alert>
      </Container>
    );
  }

  if (!course) {
    return (
      <Container>
        <Typography variant="h6" align="center" gutterBottom>
          No se encontraron detalles del curso.
        </Typography>
      </Container>
    );
  }

  return (
    <>
      <Container>
        <Card>
          <CardMedia
            component="img"
            style={{ width: 300, height: 250 }}
            image={course.imageURL}
            alt={course.nombre}
          />
          <CardContent>
            <Typography variant="h4" component="div" gutterBottom>
              {course.nombre}
            </Typography>
            <Typography variant="body1" color="text.primary">
              Categoría: {course.categoria}
            </Typography>
            <Typography variant="body1" color="text.primary">
              Dificultad: {course.dificultad}
            </Typography>
            <Typography variant="body1" color="text.primary">
              Precio: ${course.precio}
            </Typography>
            <Typography variant="body1" color="text.secondary" paragraph>
              {course.descripcion}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Creado el: {new Date(course.created_at).toLocaleDateString()}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Actualizado el: {new Date(course.updated_at).toLocaleDateString()}
            </Typography>
          </CardContent>
        </Card>
      </Container>
      <Upload />
      <Chat courseId={course.ID} />
    </>
  );
};

export default CourseDetails;
