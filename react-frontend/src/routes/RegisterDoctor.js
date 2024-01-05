import React, { useEffect } from 'react';
import RegisterDoctorComponent from '../components/auth/RegisterDoctor';
import { validateJwtToken, verifyJWTRole } from '../utils/utils';
import { JWT_COOKIE_NAME, ROLE_ADMIN } from '../utils/constants';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { LOGIN_ENDPOINT } from '../utils/endpoints';
import Cookies from 'js-cookie';

const RegisterDoctor = () => {
  const navigate = useNavigate();
  const jwtToken = Cookies.get(JWT_COOKIE_NAME);

  useEffect(() => {
    // Validate if JWT is available
    if (!jwtToken) {
      console.error('JWT token not found in cookies');
      toast.error('Error: JWT token not found. Please log in.');
      navigate(LOGIN_ENDPOINT);
      return;
    }

    const decodedClaims = validateJwtToken(jwtToken);
    const roles = [ROLE_ADMIN]
    if (!verifyJWTRole(decodedClaims, roles)) {
      toast.error('Admin authentication required');
      navigate(LOGIN_ENDPOINT);
    }
  }, [navigate, jwtToken]);


  
  return (
    <div className="container mt-4">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card">
            <div className="card-body">
              <h2 className="card-title text-center mb-4">Enter Doctor Data</h2>
              <RegisterDoctorComponent jwtToken={jwtToken} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegisterDoctor;
