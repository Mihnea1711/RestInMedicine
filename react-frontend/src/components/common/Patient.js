import React from 'react';

export const Patient = ({ patient, onSeeMore }) => {
  return (
    <li key={patient.idPatient}>
      <span>{patient.firstName}</span>
      <span style={{ marginLeft: '8px' }}>{patient.secondName}</span>
      <span style={{ marginLeft: '8px' }}>{patient.email}</span>
      <span style={{ marginLeft: '8px' }}>{patient.phoneNumber}</span>
      <span style={{ marginLeft: '8px' }}>{patient.cnp}</span>
      <span style={{ marginLeft: '8px' }}>{patient.birthDay}</span>
      <button onClick={() => onSeeMore(patient.idPatient)}>See More</button>
    </li>
  );
};
