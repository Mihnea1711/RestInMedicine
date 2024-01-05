import React from 'react';
import { ROLE_DOCTOR } from '../../utils/constants';
import { verifyJWTRole } from '../../utils/utils';
import { useNavigate } from 'react-router-dom';
import { UPDATE_APPOINTMENT_ENDPOINT } from '../../utils/endpoints';

export const Appointment = ({ appointment, claims }) => {
  const navigate = useNavigate();
  const doctorRole = [ROLE_DOCTOR];

  const handleEditAppointment = () => {
    navigate(UPDATE_APPOINTMENT_ENDPOINT(appointment.idProgramare));
  };

  const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString(undefined, options);
  };

  return (
    <tr key={appointment.idProgramare}>
      <td>{formatDate(appointment.date)}</td>
      <td>{appointment.status}</td>
      {verifyJWTRole(claims, doctorRole) && (
        <td className='d-flex justify-content-center'>
          <button className="btn btn-outline-primary btn-sm mx-2" onClick={handleEditAppointment}>Edit</button>
        </td>
      )}
    </tr>
  );
};
