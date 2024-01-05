import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { ROLE_DOCTOR } from '../../utils/constants';
import { verifyJWTRole } from '../../utils/utils';
import { UPDATE_CONSULTATION_ENDPOINT } from '../../utils/endpoints';

export const Consultation = ({ consultation, claims }) => {
  const doctorRole = [ROLE_DOCTOR];
  const navigate = useNavigate();
  const [showInvestigations, setShowInvestigations] = useState(false);

  const handleEditAppointment = () => {
    navigate(UPDATE_CONSULTATION_ENDPOINT(consultation.idConsultation));
  };

  const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString(undefined, options);
  };

  const handleToggleInvestigations = () => {
    setShowInvestigations(!showInvestigations);
  };
  
  return (
    <tr key={consultation.idConsultation}>
      <td>{formatDate(consultation.date)}</td>
      <td>{consultation.diagnostic}</td>
      <td>
        {showInvestigations ? (
          <>
            <ul>
              {consultation.investigations.map((investigation, index) => (
                <li key={index}>
                  <strong>Name:</strong> {investigation.name},{' '}
                  <strong>Processing Time:</strong> {investigation.processingTime},{' '}
                  <strong>Result:</strong> {investigation.result}
                </li>
              ))}
            </ul>
            <button className="btn btn-link" onClick={handleToggleInvestigations}>
              Hide Investigations
            </button>
          </>
        ) : (
          <button className="btn btn-link" onClick={handleToggleInvestigations}>
            Show Investigations
          </button>
        )}
      </td>
      {verifyJWTRole(claims, doctorRole) && (
        <td>
            <button className="btn btn-outline-primary btn-sm mx-2" onClick={handleEditAppointment}>
              Edit
            </button>
        </td>
      )}
    </tr>
  );
};
