// src/components/AboutUs.js
import React from "react";
import {
  Container,
  Typography,
  Grid,
  Card,
  CardContent,
  CardMedia,
  Box,
} from "@mui/material";

const teamMembers = [
  {
    name: "Felipe Ganame",
    role: "CEO & Founder",
    image:
      "https://cdn.urbantecno.com/urbantecno/s/2023-01-05-11-27-elon-musk.png", // Reemplaza con la URL de la imagen real
    description:
      "Felipe es el fundador de nuestra plataforma de cursos con más de 20 años de experiencia en la industria educativa.",
  },
  {
    name: "Joaquin Lista",
    role: "CTO",
    image:
      "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcS4Zmm2GK0bUvH0HebvXiTu3OV9AjJhXYsYxzIlSUuoSFv2cDdZ", // Reemplaza con la URL de la imagen real
    description:
      "Joaco lidera nuestro equipo de tecnología, asegurando que nuestra plataforma esté siempre a la vanguardia.",
  },
  {
    name: "Arnon Nahmias",
    role: "COO",
    image: "https://i.blogs.es/c3747a/steve-jobs-presentacion/375_375.webp",
    description:
      "Arnon es nuestro director de operaciones, con una gran experiencia en la gestión y optimización de procesos empresariales.",
  },
];

const AboutUs = () => {
  return (
    <Container maxWidth="lg">
      <Box mt={5}>
        <Typography variant="h3" align="center" gutterBottom>
          Sobre Nosotros
        </Typography>
        <Typography variant="body1" align="center" paragraph>
          Nuestra misión es proporcionar cursos de alta calidad para todos, en
          cualquier momento y en cualquier lugar. Creemos en el poder de la
          educación para transformar vidas y estamos comprometidos a brindar una
          experiencia de aprendizaje excepcional.
        </Typography>
        <Typography variant="h4" align="center" gutterBottom>
          Nuestro Equipo
        </Typography>
        <Grid container spacing={4}>
          {teamMembers.map((member, index) => (
            <Grid item key={index} xs={12} sm={6} md={4}>
              <Card>
                <CardMedia
                  component="img"
                  height="250"
                  image={member.image}
                  alt={member.name}
                />
                <CardContent>
                  <Typography variant="h5" component="div">
                    {member.name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {member.role}
                  </Typography>
                  <Typography variant="body1" color="text.primary">
                    {member.description}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
        <Box mt={5}>
          <Typography variant="h4" align="center" gutterBottom>
            Nuestra Visión
          </Typography>
          <Typography variant="body1" align="center" paragraph>
            Ser la plataforma de cursos en línea líder a nivel mundial,
            reconocida por la calidad de nuestros contenidos y la efectividad de
            nuestros métodos de enseñanza.
          </Typography>
          <Typography variant="h4" align="center" gutterBottom>
            Nuestros Valores
          </Typography>
          <Typography variant="body1" align="center" paragraph>
            Innovación, calidad, accesibilidad y compromiso con nuestros
            estudiantes. Nos esforzamos por mejorar continuamente y brindar el
            mejor servicio posible.
          </Typography>
        </Box>
      </Box>
    </Container>
  );
};

export default AboutUs;
