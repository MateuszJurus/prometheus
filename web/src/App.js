import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navigation from './components/Navigation/Navigation.js';
import Home from './pages/Home/Home.js';
import About from './pages/About/About.js';
import FormPage from './pages/Form/Form.js';

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/about', name: 'About', component: About },
  { path: '/form', name: 'Form', component: FormPage },
]

function App() {
  return (
    <Router>
      <div>
        <Navigation routes={routes} />
        <main className='main'>
        <Routes>
          {routes.map((route, index) => (
            <Route key={index} path={route.path} element={<route.component />} />
          ))}
        </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
