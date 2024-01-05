import React, { useEffect, useState } from 'react';
import { validateJwtToken, verifyJWTRole } from '../utils/utils';
import { JWT_COOKIE_NAME, ROLE_PATIENT } from '../utils/constants';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { LOGIN_ENDPOINT } from '../utils/endpoints';
import AllDoctorsComponent from '../components/main/Doctors';
import Cookies from 'js-cookie';

const AllDoctors = () => {
  const jwtToken = Cookies.get(JWT_COOKIE_NAME);

  const [claims, setClaims] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = () => {
      try {
        if (!jwtToken) {
          throw new Error('JWT token not found in cookies');
        }
        
        const decodedClaims = validateJwtToken(jwtToken);
        setClaims(decodedClaims);

        const roles = [ROLE_PATIENT];
        if (!verifyJWTRole(decodedClaims, roles)) {
          throw new Error('Patient authentication required');
        }
      } catch (error) {
        console.error('Error during authentication:', error.message);
        toast.error(`Authentication error: ${error.message}`);
        navigate(LOGIN_ENDPOINT);
      }
    }
    fetchData();
  }, [navigate, jwtToken]);
  
  return (
    <div className="container mt-4">
      <h2>Doctor List</h2>
      <AllDoctorsComponent jwt={jwtToken} claims={claims}/>
    </div>
  );
};

export default AllDoctors;
