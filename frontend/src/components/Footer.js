import React from 'react';
import { Box, Container, Grid, Typography, Link, IconButton } from '@mui/material';
import FacebookIcon from '@mui/icons-material/Facebook';
import TwitterIcon from '@mui/icons-material/X';
import InstagramIcon from '@mui/icons-material/Instagram';
import WhatsAppIcon from '@mui/icons-material/WhatsApp';
import PhoneIcon from '@mui/icons-material/Phone';

const Footer = () => {
  return (
    <Box sx={{ bgcolor: 'primary.main', color: 'white', p: 3 }} component="footer">
      <Container maxWidth="lg">
        <Grid container spacing={3} alignItems="center">
          <Grid item xs={12} md={4}>
            <Typography variant="body2">
              © My Online Courses {new Date().getFullYear()}
            </Typography>
          </Grid>
          <Grid item xs={12} md={4} textAlign="center">
            <Link href="tel:0800-0000-0000" color="inherit" underline="none">
              <PhoneIcon sx={{ verticalAlign: 'middle', mr: 1 }} /> 0800-0000-0000
            </Link>
          </Grid>
          <Grid item xs={12} md={4} textAlign="center">
            <Typography variant="body2" component="span">
              Follow us on
            </Typography>
            <IconButton href="https://www.facebook.com/" color="inherit">
              <FacebookIcon />
            </IconButton>
            <IconButton href="https://x.com/" color="inherit">
              <TwitterIcon />
            </IconButton>
            <IconButton href="https://www.instagram.com/" color="inherit">
              <InstagramIcon />
            </IconButton>
            <IconButton href="https://www.whatsapp.com/" color="inherit">
              <WhatsAppIcon />
            </IconButton>
          </Grid>
        </Grid>
      </Container>
    </Box>
  );
};

export default Footer;
