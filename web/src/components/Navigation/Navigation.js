// src/components/Navigation.js
import React from 'react';
import { useAuth } from '../../utils/AuthProvider';
import './Navigation.css';
import NavigationItem from '../NavigationItem/NavigationItem';
import LoginForm from '../LoginForm/LoginForm';

const Navigation = ({ routes }) => {
  const { isAuthenticated, logout } = useAuth();

  return (
    <nav className='navigation'>
      <ul className='navigation__list'>
        {routes.map((route, index) => (
          <NavigationItem key={index} name={route.name} url={route.path} />
        ))}
      </ul>
      {isAuthenticated ? (
        <button onClick={logout}>Logout</button> // This could be a Logout component instead
      ) : (
        <LoginForm /> // Show LoginForm if not authenticated
      )}
    </nav>
  );
};

export default Navigation;