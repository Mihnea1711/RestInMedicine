import React, {  } from 'react';
import axios from 'axios';
import { GATEWAY_UPDATE_CONSULTATION, GATEWAY_CREATE_CONSULTATION, GATEWAY_GET_DOCTOR_BY_USER_ID, CONSULTATION_HISTORY_ENDPOINT } from '../../utils/endpoints';
import { buildURL, handleResponse, parseJwt } from '../../utils/utils';
import { ConsultationData } from '../../models/Consultation';
import { format } from 'date-fns';
import { TIMESTAMP_SUFFIX } from '../../utils/constants';
import { toast } from 'react-toastify';
import ConsultationForm from '../forms/ConsultationForm';
import { useNavigate } from 'react-router-dom';

const CreateEditConsultationComponent = ({ jwt, consultation }) => {
  const navigate = useNavigate();

  const handleFormSubmit = async (formData) => {
    try {
      const headers = {
        Authorization: `Bearer ${jwt}`,
      };

      // A doctor creates consultations to himself
      const userID = parseJwt(jwt).sub;
      const doctorIDURL = buildURL(GATEWAY_GET_DOCTOR_BY_USER_ID + userID, "");
      const doctorIDResponse = await handleResponse(axios.get(doctorIDURL, { headers }));
      const doctorID = doctorIDResponse.payload.idDoctor;

      const requestData = new ConsultationData (
        consultation?.idConsultation || null,
        parseInt(formData.idPatient, 10) || consultation.idPatient,
        doctorID,
        format(formData.date, 'yyyy-MM-dd') + TIMESTAMP_SUFFIX,
        formData.diagnostic,
        formData.investigations
      );

      let request;
      if (consultation) {
        // If editing an existing consultation, use the update endpoint
        const updateConsultationURL = buildURL(GATEWAY_UPDATE_CONSULTATION + consultation.idConsultation, "");
        request = axios.put(updateConsultationURL, requestData, { headers });
      } else {
        // If creating a new consultation, use the create endpoint
        const createConsultationURL = buildURL(GATEWAY_CREATE_CONSULTATION, "");
        request = axios.post(createConsultationURL, requestData, { headers });
      }

      const responseData = await handleResponse(request);
      toast.success(responseData.message);
      navigate(CONSULTATION_HISTORY_ENDPOINT);
    } catch (error) {
      console.error('Error saving consultation:', error.message);
      toast.error(`Error saving consultation: ${error.message}`);
    }
  };

  return (
    <div className="container mt-4">
      <div className="row justify-content-center">
        <div className="col-md-8">
          <ConsultationForm onSubmit={handleFormSubmit} initialData={consultation} />
        </div>
      </div>
    </div>
  );
};

export default CreateEditConsultationComponent;