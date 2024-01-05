import React, { useEffect, useState } from 'react';
import { validateJwtToken, verifyJWTRole } from '../utils/utils';
import { JWT_COOKIE_NAME, ROLE_DOCTOR, ROLE_PATIENT } from '../utils/constants';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { CREATE_APPOINTMENT_ENDPOINT, LOGIN_ENDPOINT } from '../utils/endpoints';
import AppointmentHistoryComponent from '../components/main/AppointmentHistory';
import Cookies from 'js-cookie';

const AppointmentHistory = () => {
  const jwtToken = Cookies.get(JWT_COOKIE_NAME);

  const [claims, setClaims] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = () => {
      try {
        if (!jwtToken) {
          throw new Error('JWT token not found in cookies');
        }

        const claims = validateJwtToken(jwtToken);
        setClaims(claims);

        const allowedRoles = [ROLE_PATIENT, ROLE_DOCTOR];
        if (!verifyJWTRole(claims, allowedRoles)) {
          throw new Error('Patient or Doctor authentication required');
        }
      } catch (error) {
        console.error('Error during authentication:', error.message);
        toast.error(`Authentication error: ${error.message}`);
        navigate(LOGIN_ENDPOINT);
      }
    }

    fetchData();
  }, [navigate, jwtToken]);

  const doctorRole = [ROLE_DOCTOR];
  return (
    <div className="container mt-4">
      {claims && (
        <>
          <h2>Appointment History</h2>
          {verifyJWTRole(claims, doctorRole) && (
            <div className="mb-3">
              <a href={CREATE_APPOINTMENT_ENDPOINT} className="btn btn-primary">
                Create Appointment
              </a>
            </div>
          )}
          <AppointmentHistoryComponent jwt={jwtToken} claims={claims} />
        </>
      )}
    </div>
  );
  
};

export default AppointmentHistory;
