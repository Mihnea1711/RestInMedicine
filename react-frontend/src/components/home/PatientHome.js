import React from 'react';
import { Link } from 'react-router-dom';
import Cookies from 'js-cookie';
import { toast } from 'react-toastify';
import { JWT_COOKIE_NAME } from '../../utils/constants';
import {
    APPOINTMENT_HISTORY_ENDPOINT,
    CONSULTATION_HISTORY_ENDPOINT,
    DOCTORS_ENDPOINT,
    PROFILE_ENDPOINT
} from '../../utils/endpoints';


const PatientHomeComponent = () => {
  const handleLogout = () => {
    Cookies.remove(JWT_COOKIE_NAME);
    toast.success('Log Out successful');
    window.location.reload();
  };

  return (
    <div className="container mt-4">
    <h2>Welcome, Patient!</h2>
    <ul className="list-group">
      <li className="list-group-item">
        <Link to={PROFILE_ENDPOINT}>See Profile</Link>
      </li>
      <li className="list-group-item">
        <Link to={APPOINTMENT_HISTORY_ENDPOINT}>Appointment History</Link>
      </li>
      <li className="list-group-item">
        <Link to={CONSULTATION_HISTORY_ENDPOINT}>Consultation History</Link>
      </li>
      <li className="list-group-item">
        <Link to={DOCTORS_ENDPOINT}>See Doctors</Link>
      </li>
    </ul>
    <button className="btn btn-danger mt-3" onClick={handleLogout}>
      Logout
    </button>
  </div>
  );
};

export default PatientHomeComponent;
