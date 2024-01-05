import React from 'react';

export const Doctor = ({ doctor }) => {
  return (
    <tr key={doctor.idDoctor}>
      <td>{doctor.firstName}</td>
      <td>{doctor.secondName}</td>
      <td>{doctor.email}</td>
      <td>{doctor.phoneNumber}</td>
      <td>{doctor.specialization}</td>
    </tr>
  );
};
