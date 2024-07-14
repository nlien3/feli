import React from 'react';
import { Route, Routes } from 'react-router-dom';
import './App.css';
import Header from './components/Header';
import Footer from './components/Footer';
import HomePage from './pages/HomePage';
import AboutPage from './pages/AboutPage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import CourseDetailPage from './pages/CourseDetailPage'; // Importa el nuevo componente
import MyCourses from './pages/MyCourses';
import CourseDetails from './pages/CourseDetails';
import CreateCourse from './pages/CreateCourse';
import EditCourse from './pages/EditCourse';
function App() {
  return (
    <div className="App">
      <Header />
      <main style={{ flex: 1 }}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/about" element={<AboutPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/course/:id" element={<CourseDetailPage />} />

          <Route path="/my-courses" element={<MyCourses />} />
          <Route path="/my-courses/:id" element={<CourseDetails />} />
          <Route path="/create-course" element={<CreateCourse />} />
          <Route path="/course/edit/:id" element={<EditCourse />} />
        </Routes>
      </main>
      <Footer />
    </div>
  );
}


export default App;
