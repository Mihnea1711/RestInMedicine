import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import Cookies from 'js-cookie';
import axios from 'axios';
import { GATEWAY_GET_CONSULTATION } from '../utils/endpoints';
import { buildURL, handleResponse, validateJwtToken } from '../utils/utils';
import { verifyJWTRole } from '../utils/utils';
import { JWT_COOKIE_NAME, ROLE_DOCTOR } from '../utils/constants';
import CreateEditConsultationComponent from '../components/main/CreateEditConsultation';
import { toast } from 'react-toastify';

const CreateEditConsultation = () => {      
    const jwtToken = Cookies.get(JWT_COOKIE_NAME);
    const { consultationID } = useParams();
    const [consultation, setConsultation] = useState(null);
    const [claims, setClaims] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                if (!jwtToken) {
                    throw new Error('JWT token not found in cookies');
                }

                const claims = validateJwtToken(jwtToken);
                setClaims(claims);

                if (consultationID) {
                    const headers = {
                        Authorization: `Bearer ${jwtToken}`,
                    };
                    const consultationURL = buildURL(GATEWAY_GET_CONSULTATION, consultationID);
                    const request = axios.get(consultationURL, { headers });
                    const responseData = await handleResponse(request);
                    setConsultation(responseData.payload);
                }
            } catch (error) {
                console.error('Error during data fetching:', error.message);
                toast.error(`Error during data fetching: ${error.message}`);
            }
        };

        fetchData();
    }, [consultationID, jwtToken]);

    // Check doctor role based on JWT
    const doctorRole = [ROLE_DOCTOR];
    return (
        <div className="container mt-4">
            {claims && (
                <>
                    <h2>{consultationID ? 'Edit Consultation' : 'Create Consultation'}</h2>
                    {verifyJWTRole(claims, doctorRole) ? (
                        <CreateEditConsultationComponent jwt={jwtToken} consultation={consultation} />
                    ) : (
                        <p>Doctor authentication required.</p>
                    )}
                </>
            )}
        </div>
    );
};

export default CreateEditConsultation;
