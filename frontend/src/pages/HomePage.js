import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Box, Card, CardContent, CardMedia, Typography, Grid, CircularProgress, Alert, TextField, Button, MenuItem, Select, FormControl, InputLabel } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const HomePage = () => {
  const [courses, setCourses] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [searchType, setSearchType] = useState('name');
  const navigate = useNavigate();

  useEffect(() => {
    const getCourses = async () => {
      try {
        const response = await axios.get(process.env.REACT_APP_API + "/courses");
        setCourses(response.data);
      } catch (error) {
        console.error('Error fetching courses', error);
        setError('Error fetching courses');
      } finally {
        setLoading(false);
      }
    };

    getCourses();
  }, []);

  const handleSearch = async () => {
    if (searchQuery.trim() === '') {
      setSearchResults([]);
      return;
    }

    setLoading(true);
    setError(null);

    try {
      let combinedResults = [];
      const lowerCaseQuery = searchQuery.toLowerCase();

      if (searchType === 'id') {
        if (!isNaN(searchQuery)) {
          try {
            const responseById = await axios.get(`${process.env.REACT_APP_API}/courses/${searchQuery}`);
            if (responseById.data) {
              combinedResults = [responseById.data];
            }
          } catch (err) {
            console.log(`ID search failed: ${err.message}`);
          }
        }
      } else if (searchType === 'name') {
        try {
          const responseByName = await axios.get(`${process.env.REACT_APP_API}/courses/name/${lowerCaseQuery}`);
          if (responseByName.data && responseByName.data.length > 0) {
            combinedResults = responseByName.data;
          }
        } catch (err) {
          console.log(`Name search failed: ${err.message}`);
        }
      } else if (searchType === 'category') {
        try {
          const responseByCategory = await axios.get(`${process.env.REACT_APP_API}/courses/category/${lowerCaseQuery}`);
          if (responseByCategory.data && responseByCategory.data.length > 0) {
            combinedResults = responseByCategory.data;
          }
        } catch (err) {
          console.log(`Category search failed: ${err.message}`);
        }
      }

      if (combinedResults.length === 0) {
        throw new Error('No courses found');
      }

      setSearchResults(combinedResults);
    } catch (error) {
      console.error('Error fetching courses:', error.response?.data || error.message);
      setError(`Error fetching courses: ${error.response?.data?.error || error.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handleCourseClick = (id) => {
    navigate(`/course/${id}`);
  };

  const styles = {
    courseCard: {
      cursor: 'pointer',
    },
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Alert severity="error">{error}</Alert>
      </Box>
    );
  }

  const displayedCourses = searchResults.length > 0 ? searchResults : courses;

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Available Courses
      </Typography>
      <Box display="flex" alignItems="center" mb={2}>
        <FormControl variant="outlined" sx={{ minWidth: 120, marginRight: 2 }}>
          <InputLabel id="search-type-label">Search By</InputLabel>
          <Select
            labelId="search-type-label"
            value={searchType}
            onChange={(e) => setSearchType(e.target.value)}
            label="Search By"
          >
            <MenuItem value="id">ID</MenuItem>
            <MenuItem value="name">Name</MenuItem>
            <MenuItem value="category">Category</MenuItem>
          </Select>
        </FormControl>
        <TextField
          label="Search Courses"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          variant="outlined"
          fullWidth
          margin="normal"
        />
        <Button
          variant="contained"
          color="primary"
          onClick={handleSearch}
          style={{ marginLeft: "16px", height: "56px" }}
        >
          Search
        </Button>
      </Box>
      <Grid container spacing={3}>
        {displayedCourses.map((course) => (
          <Grid item xs={12} sm={6} md={4} key={course.ID}>
            <Card
              style={styles.courseCard}
              onClick={() => handleCourseClick(course.ID)}
            >
              <CardMedia
                component="img"
                style={{ width: 300, height: 250 }}
                image={course.imageURL}
                alt={course.nombre}
              />
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {course.nombre}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Category: {course.categoria}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Price: ${course.precio}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default HomePage;
