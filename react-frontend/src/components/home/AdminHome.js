
import React from 'react';
import { Link } from 'react-router-dom';
import { toast } from 'react-toastify';
import Cookies from 'js-cookie';
import { JWT_COOKIE_NAME } from '../../utils/constants';
import { REGISTER_DOCTOR_ENDPOINT, PROFILE_ENDPOINT, USERS_ENDPOINT } from '../../utils/endpoints';

const AdminHomeComponent = () => {
  const handleLogout = () => {
    Cookies.remove(JWT_COOKIE_NAME);
    toast.success('Log Out successful');
    window.location.reload();
  };

  return (
    <div className="container mt-4">
      <h2>Welcome, Admin!</h2>
      <ul className="list-group">
        <li className="list-group-item">
          <Link to={PROFILE_ENDPOINT}>Profile</Link>
        </li>
        <li className="list-group-item">
          <Link to={REGISTER_DOCTOR_ENDPOINT}>Register Doctor</Link>
        </li>
        <li className="list-group-item">
          <Link to={USERS_ENDPOINT}>See All Users</Link>
        </li>
      </ul>
      <button className="btn btn-danger mt-3" onClick={handleLogout}>
        Logout
      </button>
    </div>
  );
};

export default AdminHomeComponent;
