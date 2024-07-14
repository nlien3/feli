import React, { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { Box, Typography, CircularProgress, Alert, Card, CardMedia, CardContent, Button } from '@mui/material';
import AuthContext from '../context/AuthContext';
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const CourseDetailPage = () => {
  const { id } = useParams();
  const navigate = useNavigate(); 
  const [course, setCourse] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [subscribeError, setSubscribeError] = useState(null);
  const [subscribeSuccess, setSubscribeSuccess] = useState(null);
  const { user } = useContext(AuthContext);

  const handleDelete = async () => {
    try {
      const response = await axios.delete(
        `${process.env.REACT_APP_API}/courses/${id}`,
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        }
      );
      alert("Course successfully deleted");
      navigate("/");
    } catch (err) {
      console.error("Error deleting the course", err);
      alert("Failed to delete the course");
    }
  };

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await axios.get(
          `${process.env.REACT_APP_API}/courses/${id}`
        );
        setCourse(response.data);
      } catch (error) {
        console.error("Error fetching course details", error);
        setError("Error fetching course details");
      } finally {
        setLoading(false);
      }
    };

    fetchCourse();
  }, [id]);

  const handleSubscribe = async () => {
    console.log("User:", user.userId); // Add this line for debugging
    if (!user || !user.userId) {
      setSubscribeError("You must be logged in to subscribe");
      return;
    }

    try {
      const response = await axios.post(
        process.env.REACT_APP_API + "/subscriptions",
        {
          userID: parseInt(user.userId),
          courseID: parseInt(id),
        },
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        }
      );

      setSubscribeSuccess("Subscribed successfully!");
    } catch (err) {
      console.error("Error subscribing to the course", err);
      if (err?.response?.data?.details == "la suscripción ya existe") {
        setSubscribeError("You are already subscribe to this course");
      } else {
        setSubscribeError("Error subscribing to the course");
      }
    }
  };

  if (loading) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <Alert severity="error">{error}</Alert>
      </Box>
    );
  }

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Card>
        {course?.imageURL && (
          <Box sx={{ maxWidth: 500, maxHeight: 500, mx: "auto" }}>
            <CardMedia
              component="img"
              image={course.imageURL}
              alt={course.nombre}
              sx={{ maxWidth: 500, maxHeight: 500 }}
            />
          </Box>
        )}
        <CardContent>
          <Typography variant="h4" component="h1" gutterBottom>
            {course.nombre}
          </Typography>
          {user?.userType === "normal" && (
            <Button
              variant="contained"
              color="primary"
              onClick={handleSubscribe}
              sx={{ mb: 2 }}
            >
              Subscribe to this Course
            </Button>
          )}
          {user?.userType === "admin" && (
            <>
              <Button
                component={Link}
                to={`/my-courses/${course.ID}`}
                variant="contained"
                color="primary"
              >
                Más Información
              </Button>
              <Button
                variant="contained"
                sx={{ backgroundColor: "#FFA500", color: "black", ml: 2 }} // Naranja personalizado
                component={Link}
                to={`/course/edit/${course.ID}`}
              >
                Editar
              </Button>
              <Button
                variant="contained"
                color="error"
                onClick={handleDelete}
                sx={{ ml: 2 }}
              >
                Delete
              </Button>
            </>
          )}
          {subscribeError && <Alert severity="error">{subscribeError}</Alert>}
          {subscribeSuccess && (
            <Alert severity="success">{subscribeSuccess}</Alert>
          )}
          <Typography variant="body1" paragraph>
            {course.descripcion}
          </Typography>
          <Typography variant="body1" paragraph>
            Category: {course.categoria}
          </Typography>
          <Typography variant="body1" paragraph>
            Difficulty: {course.dificultad}
          </Typography>
          <Typography variant="body1" paragraph>
            Price: ${course.precio}
          </Typography>
        </CardContent>
      </Card>
    </Box>
  );
};

export default CourseDetailPage;
