import React, { createContext, useContext, useEffect, useState } from 'react';

const AuthContext = createContext();

export function useAuth() {
  return useContext(AuthContext);
}

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(!localStorage.getItem('pr-token'));

  useEffect(() => {
    console.log(isAuthenticated)
    },[isAuthenticated])

useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
        // Validate token. If invalid, call logout()
        logout();
    }
    }, []);

  const login = async (username, password) => {
    try {
        const response = await fetch('http://localhost:8080/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username, password }),
        });
    
        if (!response.ok) {
          throw new Error('Login failed');
        }
    
        const { token } = await response.json();
        localStorage.setItem('pr-token', token); // Store the token for future requests
        setIsAuthenticated(true); // Update the authentication state
      } catch (error) {
        console.error(error);
      }
  };

  const logout = () => {
    setIsAuthenticated(false);
    // On logout, set isAuthenticated to false
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}
