import React from 'react';
import { Link } from 'react-router-dom';
import Cookies from 'js-cookie';
import { toast } from 'react-toastify';
import { JWT_COOKIE_NAME } from '../../utils/constants';
import {
  REGISTER_PATIENT_ENDPOINT,
  PROFILE_ENDPOINT,
  APPOINTMENT_HISTORY_ENDPOINT,
  CONSULTATION_HISTORY_ENDPOINT
} from '../../utils/endpoints';


const DoctorHomeComponent = () => {
  const handleLogout = () => {
    Cookies.remove(JWT_COOKIE_NAME);
    toast.success('Log Out successful');
    window.location.reload();
  };

  return (
    <div className="container mt-4">
    <h2>Welcome, Doctor!</h2>
    <ul className="list-group">
      <li className="list-group-item">
        <Link to={PROFILE_ENDPOINT}>See Profile</Link>
      </li>
      <li className="list-group-item">
        <Link to={REGISTER_PATIENT_ENDPOINT}>Register Patient</Link>
      </li>
      <li className="list-group-item">
        <Link to={APPOINTMENT_HISTORY_ENDPOINT}>See Appointments</Link>
      </li>
      <li className="list-group-item">
        <Link to={CONSULTATION_HISTORY_ENDPOINT}>See Consultations</Link>
      </li>
    </ul>
    <button className="btn btn-danger mt-3" onClick={handleLogout}>
      Logout
    </button>
  </div>
  );
};

export default DoctorHomeComponent;
