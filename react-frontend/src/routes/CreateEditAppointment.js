import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import Cookies from 'js-cookie';
import axios from 'axios';
import { GATEWAY_GET_APPOINTMENT } from '../utils/endpoints';
import { buildURL, handleResponse, validateJwtToken } from '../utils/utils';
import { verifyJWTRole } from '../utils/utils';
import { JWT_COOKIE_NAME, ROLE_DOCTOR } from '../utils/constants';
import CreateEditAppointmentComponent from '../components/main/CreateEditAppointment';
import { toast } from 'react-toastify';

const CreateEditAppointment = () => {
    const jwtToken = Cookies.get(JWT_COOKIE_NAME);
    const { appointmentID } = useParams();
    const [appointment, setAppointment] = useState(null);
    const [claims, setClaims] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                if (!jwtToken) {
                    throw new Error('JWT token not found in cookies');
                }

                const claims = validateJwtToken(jwtToken);
                setClaims(claims);

                if (appointmentID) {
                    const headers = {
                        Authorization: `Bearer ${jwtToken}`,
                    };

                    const appointmentURL = buildURL(GATEWAY_GET_APPOINTMENT, appointmentID);
                    const request = axios.get(appointmentURL, { headers });
                    const responseData = await handleResponse(request);

                    setAppointment(responseData.payload);
                }
            } catch (error) {
                console.error('Error during data fetching:', error.message);
                toast.error(`Error during data fetching: ${error.message}`);
            }
        };

        fetchData();
    }, [appointmentID, jwtToken]);

    // Check doctor role based on JWT
    const doctorRole = [ROLE_DOCTOR];
    return (
        <div className="container mt-4">
            {claims && (
                <>
                    <h2>{appointmentID ? 'Edit Appointment' : 'Create Appointment'}</h2>
                    {verifyJWTRole(claims, doctorRole) ? (
                        <CreateEditAppointmentComponent jwt={jwtToken} appointment={appointment} />
                    ) : (
                        <p>Doctor authentication required.</p>
                    )}
                </>
            )}
        </div>
    );
};

export default CreateEditAppointment;
