import React, { useEffect, useMemo, useState } from 'react';
import axios from 'axios';
import { GATEWAY_GET_DOCTORS } from '../../utils/endpoints';
import { toast } from 'react-toastify';
import { buildURL, handleResponse } from '../../utils/utils';
import { useNavigate } from 'react-router-dom';
import { Doctor } from '../common/Doctor';
import { Spinner, Table } from 'react-bootstrap';

const AllPatientsComponent = ({claims, jwt}) => {
  const headers = useMemo(() => ({
    Authorization: `Bearer ${jwt}`,
  }), [jwt]);

  const [doctors, setDoctors] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchDoctors = async () => {
      try {
        const allDoctorsURL = buildURL(GATEWAY_GET_DOCTORS, "");
        const request = axios.get(allDoctorsURL, {headers});
        const responseData = await handleResponse(request);

        setDoctors(responseData.payload);
        setLoading(false);
        toast.success("Doctors loaded successfully.");
      } catch (error) {
        console.error('Error fetching doctors:', error.message);
        setLoading(false);
        toast.error('Error fetching doctors. Please try again.');
      }
    };

    fetchDoctors();
  }, [jwt, navigate, headers]);

  return (
    <div className="container mt-4">
      {loading ? (
        <Spinner animation="border" role="status" className="mx-auto"></Spinner>
      ) : (
        <Table striped bordered hover responsive className='text-center'>
          <thead>
            <tr>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Email</th>
              <th>Phone Number</th>
              <th>Specialization</th>
            </tr>
          </thead>
          <tbody>
            {doctors.map((doctor) => (
              <Doctor key={doctor.idDoctor} doctor={doctor} />
            ))}
      </tbody>
        </Table>
      )
      }
    </div>
  );
};

export default AllPatientsComponent;
