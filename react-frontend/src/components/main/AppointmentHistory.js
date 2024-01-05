import React, { useEffect, useMemo, useState } from 'react';
import axios from 'axios';
import { GATEWAY_GET_APPOINTMENTS } from '../../utils/endpoints';
import { buildURL, handleResponse, verifyJWTRole } from '../../utils/utils';
import { useNavigate } from 'react-router-dom';
import { Appointment } from '../common/Appointment';
import { toast } from 'react-toastify';
import { Spinner, Table } from 'react-bootstrap';
import { ROLE_DOCTOR } from '../../utils/constants';

const AppointmentHistoryComponent = ({claims, jwt}) => {
  const headers = useMemo(() => ({
    Authorization: `Bearer ${jwt}`,
  }), [jwt]);

  const [appointments, setAppointments] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAppointments = async () => {
      try {
        const allAppointmentsURL = buildURL(GATEWAY_GET_APPOINTMENTS, "");
        const request = axios.get(allAppointmentsURL, {headers});
        const responseData = await handleResponse(request);

        setAppointments(responseData.payload);
        setLoading(false);
        toast.success("Appointments loaded successfully.");
      } catch (error) {
        console.error('Error fetching appointments:', error.message);
        setLoading(false);
        toast.error("Error loading appointments. Please try again.")
      }
    };

    fetchAppointments();
  }, [jwt, navigate, headers]);

  return (
    <div className="container mt-4">
      {loading ? (
        <Spinner animation="border" role="status" className="mx-auto"></Spinner>
      ) : (
        <>
          {appointments && appointments.length > 0 ? (
            <Table striped bordered hover responsive className='text-center'>
              <thead>
                <tr>
                  <th>Date</th>
                  <th>Status</th>
                  {/* Add additional header columns as needed */}
                  {verifyJWTRole(claims, [ROLE_DOCTOR]) && <th>Actions</th>}
                </tr>
              </thead>
              <tbody>
                {appointments.map((appointment) => (
                  <Appointment
                    key={appointment.idProgramare}
                    appointment={appointment}
                    claims={claims}
                  />
                ))}
              </tbody>
            </Table>
          ) : (
            <p>No appointments found.</p>
          )}
        </>
      )}
    </div>
  );  
};

export default AppointmentHistoryComponent;
