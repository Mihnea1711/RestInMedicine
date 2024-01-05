import React, { useEffect, useState } from 'react';
import { validateJwtToken, verifyJWTRole } from '../utils/utils';
import { JWT_COOKIE_NAME, ROLE_ADMIN } from '../utils/constants';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { LOGIN_ENDPOINT } from '../utils/endpoints';
import AllUsersComponent from '../components/main/Users';
import Cookies from 'js-cookie';

const AllUsers = () => {
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

        const roles = [ROLE_ADMIN];
        if (!verifyJWTRole(decodedClaims, roles)) {
          throw new Error('Admin authentication required');
        }
      } catch (error) {
        console.error('Error during authentication:', error.message);
        toast.error(`Authentication error: ${error.message}`);
        navigate(LOGIN_ENDPOINT);
      }
    };

    fetchData();
  }, [navigate, jwtToken]);
  
  return (
    <div className="container mt-4">
      {claims && (
        <>
          <h2>User List</h2>
          <AllUsersComponent jwt={jwtToken} claims={claims}/>
        </>
      )}
    </div>
  );
  
};

export default AllUsers;
