import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import { format } from 'date-fns';
import 'react-datepicker/dist/react-datepicker.css';

import Input from '../common/Input';
import Select from '../common/Select';

import { buildURL, generateBirthdayFromCNP, handleResponse } from '../../utils/utils';
import { ROLE_PATIENT, TIMESTAMP_SUFFIX } from '../../utils/constants';
import { GATEWAY_REGISTER_PATIENT, GATEWAY_REGISTER_USER, HOME_ENDPOINT } from '../../utils/endpoints';
import { isValidCNP, validatePatientData } from '../../validation/validatePatientData';
import { RegisterData } from '../../models/RegisterData';
import { PatientData } from '../../models/Patient'
import MyDatePicker from '../common/MyDatePicker';

const RegisterPatientComponent = ({jwtToken}) => {
  const navigate = useNavigate(); // Hook for navigation
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    role: ROLE_PATIENT,
    firstName: '',
    secondName: '',
    email: '',
    phoneNumber: '',
    cnp: '',
    birthDay: '',
    isActive: true,
  });

  const headers = {
    Authorization: `Bearer ${jwtToken}`,
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    // Update the formData state
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
      // Automatically update the birthDay field if the entered value is a valid CNP
      birthDay: isValidCNP(value) ? generateBirthdayFromCNP(value) : prevData.birthDay,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    console.log(formData);

    // Basic form validation
    if (!validatePatientData(formData)) {
      return;
    }

    setIsSubmitting(true);
    try {
      // Send POST request to the server
      var patientUserReg = new RegisterData(formData.username, formData.password, formData.role);

      const registerUserURL = buildURL(GATEWAY_REGISTER_USER, "");
      const userRegRequest = axios.post(registerUserURL, patientUserReg, { headers });
      const responseDataUser = await handleResponse(userRegRequest);

      const patientUserID = responseDataUser.payload;
      const formattedBirthDay = format(formData.birthDay, 'yyyy-MM-dd') + TIMESTAMP_SUFFIX;
      const patientData = new PatientData(null, patientUserID, formData.firstName, formData.secondName, formData.email, formData.phoneNumber, formData.cnp, formattedBirthDay, formData.isActive);

      const registerPatientURL = buildURL(GATEWAY_REGISTER_PATIENT, "");
      const patientRegRequest = axios.post(registerPatientURL, patientData, { headers });
      const responseDataPatient = await handleResponse(patientRegRequest);

      console.log((responseDataPatient).message);
      toast.success('Registration successfull')
      navigate(HOME_ENDPOINT);
    } catch (error) {
      console.error('Patient registration failed:', error.message);
      toast.error("Error during registration");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <Input
        label="Username"
        type="text"
        name="username"
        value={formData.username}
        onChange={handleInputChange}
        required
      />
      <Input
        label="Password"
        type="password"
        name="password"
        value={formData.password}
        onChange={handleInputChange}
        required
      />
      <Input
        label="First Name"
        type="text"
        name="firstName"
        value={formData.firstName}
        onChange={handleInputChange}
        required
      />
      <Input
        label="Second Name"
        type="text"
        name="secondName"
        value={formData.secondName}
        onChange={handleInputChange}
        required
      />
      <Input
        label="Email"
        type="email"
        name="email"
        value={formData.email}
        onChange={handleInputChange}
        required
      />
      <Input
        label="Phone Number"
        type="tel"
        name="phoneNumber"
        value={formData.phoneNumber}
        onChange={handleInputChange}
        required
      />
      <Input
        label="CNP"
        type="text"
        name="cnp"
        value={formData.cnp}
        onChange={handleInputChange}
        required
      />
      <MyDatePicker
        disabled={true}
        selected={formData.birthDay}
        required
      />
      <Select
        label="Is Active"
        name="isActive"
        value={formData.isActive}
        options={[
          { label: 'Yes', value: true },
          { label: 'No', value: false },
        ]}
        onChange={handleInputChange}
        required
      />
      <button type="submit" disabled={isSubmitting} className="btn btn-primary btn-block mt-4">Register</button>
    </form>
  );
};

export default RegisterPatientComponent;
