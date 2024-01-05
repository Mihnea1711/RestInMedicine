import React from 'react';
import { Link } from 'react-router-dom';
import { LOGIN_ENDPOINT, REGISTER_ADMIN_ENDPOINT } from '../../utils/endpoints';

const DefaultHomeComponent = () => {
  return (
    <div className="container mt-4">
      <h2>Welcome to NoPainNoGain</h2>
      <ul className="list-group">
        <li className="list-group-item">
          <Link to={REGISTER_ADMIN_ENDPOINT}>Register Admin</Link>
        </li>
        <li className="list-group-item">
          <Link to={LOGIN_ENDPOINT}>Login</Link>
        </li>
      </ul>
    </div>
  );
};

export default DefaultHomeComponent;
