import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';

import { JWT_COOKIE_NAME } from '../utils/constants';
import { validateJwtToken } from '../utils/utils';
import AdminProfileComponent from '../components/profile/AdminProfile';
import PatientProfileComponent from '../components/profile/PatientProfile';
import DoctorProfileComponent from '../components/profile/DoctorProfile';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { LOGIN_ENDPOINT } from '../utils/endpoints';

const Profile = () => {
    const [claims, setClaims] = useState(null);
    const navigate = useNavigate();
    const jwtToken = Cookies.get(JWT_COOKIE_NAME);

    useEffect(() => {
        const fetchClaims = async () => {
            try {
              if (jwtToken) {
                const decodedClaims = validateJwtToken(jwtToken);
                setClaims(decodedClaims);
              } else {
                throw new Error('JWT token not found in cookies');
              }
            } catch (error) {
              console.error(error.message);
              toast.error(`Error: ${error.message}. Please log in.`);
              navigate(LOGIN_ENDPOINT);
            }
          };
        fetchClaims();
    }, [navigate, jwtToken]);

    return (
        <div className="container mt-4 text-center">
          {claims && (
            <>
              {claims.role === 'admin' && (
                <AdminProfileComponent userID={claims.sub} jwtToken={jwtToken} />
              )}
              {claims.role === 'patient' && (
                <PatientProfileComponent userID={claims.sub} jwtToken={jwtToken} />
              )}
              {claims.role === 'doctor' && (
                <DoctorProfileComponent userID={claims.sub} jwtToken={jwtToken} />
              )}
            </>
          )}
        </div>
      );
      
}

export default Profile;