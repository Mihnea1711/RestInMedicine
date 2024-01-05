import React, { useEffect, useMemo, useState } from 'react';
import axios from 'axios';
import { GATEWAY_GET_CONSULTATIONS } from '../../utils/endpoints';
import { ROLE_DOCTOR } from '../../utils/constants';
import { toast } from 'react-toastify';
import { buildURL, handleResponse, verifyJWTRole } from '../../utils/utils';
import { useNavigate } from 'react-router-dom';
import { Consultation } from '../common/Consultation';
import { Spinner, Table } from 'react-bootstrap';

const ConsultationHistoryComponent = ({claims, jwt}) => {
  const headers = useMemo(() => ({
    Authorization: `Bearer ${jwt}`,
  }), [jwt]);

  const [consultations, setConsultations] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchConsultations = async () => {
      try {
        const allConsultationsURL = buildURL(GATEWAY_GET_CONSULTATIONS, "");
        const request = axios.get(allConsultationsURL, {headers});
        const responseData = await handleResponse(request);

        setConsultations(responseData.payload);
        setLoading(false);
        toast.success("Consultations loaded successfully.");
      } catch (error) {
        console.error('Error fetching consultations:', error.message);
        setLoading(false);
        toast.error("Error loading consultations. Please try again.")
      }
    };

    fetchConsultations();
  }, [jwt, navigate, headers]);

  return (
    <div className="container mt-4">
      {loading ? (
        <Spinner animation="border" role="status" className="mx-auto"></Spinner>
      ) : (
        <>
          {consultations && consultations.length > 0 ? (
            <Table striped bordered hover responsive className='text-center'>
            <thead>
              <tr>
                <th>Date</th>
                <th>Status</th>
                <th>Investigations</th>
                {verifyJWTRole(claims, [ROLE_DOCTOR]) && <th>Actions</th>}
              </tr>
            </thead>
            <tbody>
              {consultations.map((consultation) => (
                <Consultation
                  key={consultation.idConsultation}
                  consultation={consultation}
                  claims={claims}
                />
              ))}
            </tbody>
          </Table>
          ) : (
            <p>No consultations found.</p>
          )}
        </>
        )
      }
    </div>
  );
};

export default ConsultationHistoryComponent;
