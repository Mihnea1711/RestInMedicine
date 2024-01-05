import React, {  } from 'react';
import axios from 'axios';
import { GATEWAY_UPDATE_APPOINTMENT, GATEWAY_CREATE_APPOINTMENT, GATEWAY_GET_DOCTOR_BY_USER_ID, APPOINTMENT_HISTORY_ENDPOINT } from '../../utils/endpoints';
import { buildURL, handleResponse, parseJwt } from '../../utils/utils';
import AppointmentForm from '../forms/AppointmentForm';
import { AppointmentData } from '../../models/Appointment';
import { format } from 'date-fns';
import { TIMESTAMP_SUFFIX } from '../../utils/constants';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';

const CreateEditAppointmentComponent = ({ jwt, appointment }) => {
  const navigate = useNavigate();

  const handleFormSubmit = async (formData) => {
    try {
      const headers = {
        Authorization: `Bearer ${jwt}`,
      };

      // A doctor creates appointments to himself
      // Get the doctor ID based on the user ID from the JWT
      const userID = parseJwt(jwt).sub;
      const doctorIDURL = buildURL(GATEWAY_GET_DOCTOR_BY_USER_ID + userID, "");
      const doctorIDResponse = await handleResponse(axios.get(doctorIDURL, { headers }));
      const doctorID = doctorIDResponse.payload.idDoctor;

      const requestData = new AppointmentData (
        appointment?.idProgramare || null,
        parseInt(formData.idPatient, 10) || appointment.idPatient,
        doctorID,
        format(formData.date, 'yyyy-MM-dd') + TIMESTAMP_SUFFIX,
        formData.status
      );

      let request;
      if (appointment) {
        // If editing an existing appointment, use the update endpoint
        const updateAppointmentURL = buildURL(GATEWAY_UPDATE_APPOINTMENT + appointment.idProgramare, "");
        request = axios.put(updateAppointmentURL, requestData, { headers });
      } else {
        // If creating a new appointment, use the create endpoint
        const createAppointmentURL = buildURL(GATEWAY_CREATE_APPOINTMENT, "");
        request = axios.post(createAppointmentURL, requestData, { headers });
      }

      const responseData = await handleResponse(request);
      toast.success(responseData.message);
      navigate(APPOINTMENT_HISTORY_ENDPOINT);
    } catch (error) {
      console.error('Error saving appointment:', error.message);
      toast.error(`Error saving appointment: ${error.message}`);
    }
  };

  return (
    <div className="container mt-4">
      <div className="row justify-content-center">
        <div className="col-md-8">
          <AppointmentForm onSubmit={handleFormSubmit} initialData={appointment} />
        </div>
      </div>
    </div>
  );
};

export default CreateEditAppointmentComponent;