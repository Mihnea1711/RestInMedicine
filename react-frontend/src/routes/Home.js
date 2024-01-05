import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import { Spinner } from 'react-bootstrap';
import DefaultHomeComponent from '../components/home/DefaultHome';
import AdminHomeComponent from '../components/home/AdminHome';
import PatientHomeComponent from '../components/home/PatientHome';
import DoctorHomeComponent from '../components/home/DoctorHome';
import { validateJwtToken } from '../utils/utils';
import { LOGIN_ENDPOINT } from '../utils/endpoints';
import { ROLE_ADMIN, ROLE_DOCTOR, ROLE_PATIENT, JWT_COOKIE_NAME } from '../utils/constants';

const Home = () => {
  const jwtToken = Cookies.get(JWT_COOKIE_NAME);
  const [claims, setClaims] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchClaims = () => {
      try {
        if (jwtToken) {
          setClaims(validateJwtToken(jwtToken));
        } else {
          setIsLoading(false); // Set loading to false immediately
        }
      } catch (error) {
        console.error('Error decoding JWT:', error.message);
        toast.error('Error decoding JWT. Please log in again.');
        navigate(LOGIN_ENDPOINT);
      } finally {
        setIsLoading(false);
      }
    };

    fetchClaims();
  }, [navigate, jwtToken]);

  return (
    <div className="container mt-5 text-center">
      {isLoading ? (
        <Spinner animation="border" role="status" className="mx-auto"></Spinner>
      ) : (
        <>
          {claims && claims.role === ROLE_ADMIN && <AdminHomeComponent />}
          {claims && claims.role === ROLE_PATIENT && <PatientHomeComponent />}
          {claims && claims.role === ROLE_DOCTOR && <DoctorHomeComponent />}
          {!claims && <DefaultHomeComponent />}
        </>
      )}
    </div>
  );
};

export default Home;
