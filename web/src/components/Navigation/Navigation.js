// src/components/Navigation.js

import React from 'react';
import './Navigation.css';
import NavigationItem from '../NavigationItem/NavigationItem';
import LoginForm from '../LoginForm/LoginForm';

const Navigation = ({ routes }) => {
  return (
    <nav className='navigation'>
      <ul className='navigation__list'>
        {routes.map((route, index) => (
          <NavigationItem key={index} name={route.name} url={route.path} />
        ))}
      </ul>
      <LoginForm />
    </nav>
  );
};

export default Navigation;