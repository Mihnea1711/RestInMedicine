import React, { useEffect, useState } from 'react';
import ChangePasswordComponent from '../components/main/ChangePassword';
import Cookies from 'js-cookie';
import { JWT_COOKIE_NAME } from '../utils/constants';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { LOGIN_ENDPOINT } from '../utils/endpoints';
import { validateJwtToken } from '../utils/utils';

const ChangePassword = () => {
  const [claims, setClaims] = useState(null);
  const [jwt, setJWT] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAndSetClaims = () => {
      try {
        const jwtToken = Cookies.get(JWT_COOKIE_NAME);
  
        // Validate if JWT is available
        if (!jwtToken) {
          console.error('JWT token not found in cookies');
          toast.error('Error: JWT token not found. Please log in.');
          navigate(LOGIN_ENDPOINT);
          return;
        }
  
        setJWT(jwtToken);
        const decodedClaims = validateJwtToken(jwtToken);
        setClaims(decodedClaims);
      } catch (error) {
        console.error('Error decoding JWT:', error.message);
        toast.error('Token error. Please log in again.');
        navigate(LOGIN_ENDPOINT);
      }
    };
    fetchAndSetClaims();
  }, [navigate]);

  return (
    <div className="container mt-4">
      {claims ? (
        <>
          <h2 className="mb-4">Change Your Password</h2>
          <ChangePasswordComponent userID={claims.sub} jwt={jwt} />
        </>
      ) : (
        <p>Error loading claims. Please log in again.</p>
      )}
    </div>
  );
};

export default ChangePassword;
