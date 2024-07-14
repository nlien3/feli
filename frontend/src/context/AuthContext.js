import React, { createContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('token');
    const userType = localStorage.getItem('userType');
    const userId = localStorage.getItem('userId');
    if (token) {
      setUser({ token, userType, userId });
    }
  }, []);

  const login = (token, userType, userId) => {
    localStorage.setItem('token', token);
    localStorage.setItem('userType', userType);
    localStorage.setItem('userId', userId);
    setUser({ token, userType, userId });
    navigate('/');
  };

  const logout = async () => {
    // Realiza una solicitud de logout al servidor para eliminar la cookie
    await fetch('/logout', { method: 'POST' });

    // Elimina el token del almacenamiento local
    localStorage.removeItem('token');
    localStorage.removeItem('userType');
    localStorage.removeItem('userId');
    setUser(null);
    navigate('/login');
  };

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
