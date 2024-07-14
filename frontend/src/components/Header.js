import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import AuthContext from '../context/AuthContext';
import logo from '../assets/My-online-course.svg';

const Header = () => {
  const { user, logout } = useContext(AuthContext);
  const headerStyles = {
    appBar: {
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
      padding: '0 20px',
      background: 'linear-gradient(to right, #2196F3, #21CBF3)',
      height: '120px', // Aquí puedes definir la altura del header
    },
    logo: {
      height: '200px', // Define la altura máxima del logo
    },
    navButtons: {
      display: 'flex',
      alignItems: 'center',
      height: '100%',
      fontSize: '20px',
    },
    button: {
      marginLeft: '16px',
      color: '#fff',
      fontSize: '1rem',
      fontWeight: 'bold',
      textDecoration: 'none',
      background: 'none',
      border: 'none',
      cursor: 'pointer',
    },
    linkButton: {
      color: '#fff',
      textDecoration: 'none',
      marginLeft: '16px',
    }
  };

  return (
    <header style={headerStyles.appBar}>
      <div>
        <Link to="/">
          <img src={logo} alt="My Online Course" style={headerStyles.logo} />
        </Link>
      </div>
      <div style={headerStyles.navButtons}>
        <Link to="/" style={headerStyles.linkButton}>Home</Link>
        {!user ? (
          <>
            <Link to="/about" style={headerStyles.linkButton}>About Us</Link>
            <Link to="/login" style={headerStyles.linkButton}>Log In</Link>
            <Link to="/register" style={headerStyles.linkButton}>Register</Link>
          </>
        ) : user.userType === 'admin' ? (
          <>
            <Link to="/create-course" style={headerStyles.linkButton}>Create Course</Link>
            <Link to="/about" style={headerStyles.linkButton}>About Us</Link>
            <button onClick={logout} style={headerStyles.button}>Log Out</button>
          </>
        ) : (
          <>
            <Link to="/my-courses" style={headerStyles.linkButton}>My Courses</Link>
            <Link to="/about" style={headerStyles.linkButton}>About Us</Link>
            <button onClick={logout} style={headerStyles.button}>Log Out</button>
          </>
        )}
      </div>
    </header>
  );
};

export default Header;
