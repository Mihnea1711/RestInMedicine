import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';

import Input from '../common/Input';
import Select from '../common/Select';

import { buildURL, handleResponse } from '../../utils/utils';
import { ROLE_DOCTOR, ValidSpecializations } from '../../utils/constants';
import { GATEWAY_REGISTER_DOCTOR, GATEWAY_REGISTER_USER, HOME_ENDPOINT } from '../../utils/endpoints';
import { validateDoctorData } from '../../validation/validateDoctorData';
import { RegisterData } from '../../models/RegisterData';
import { DoctorData } from '../../models/Doctor';

const RegisterDoctorComponent = ({jwtToken}) => {
  const navigate = useNavigate();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    role: ROLE_DOCTOR,
    firstName: '',
    secondName: '',
    email: '',
    phoneNumber: '',
    specialization: '',
    isActive: true,
  });

  const headers = {
    Authorization: `Bearer ${jwtToken}`,
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic form validation
    if (!validateDoctorData(formData)) {
      return;
    }

    setIsSubmitting(true);
    try {
      // Send POST request to the server
      var doctorUserReg = new RegisterData(formData.username, formData.password, formData.role);
      const registerUserURL = buildURL(GATEWAY_REGISTER_USER, "");
      const userRegRequest = axios.post(registerUserURL, doctorUserReg, { headers });
      const responseDataUser = await handleResponse(userRegRequest);
      const doctorUserID = responseDataUser.payload;

      const doctorData = new DoctorData(null, doctorUserID, formData.firstName, formData.secondName, formData.email, formData.phoneNumber, formData.specialization, formData.isActive);
      const registerDoctorURL = buildURL(GATEWAY_REGISTER_DOCTOR, "");
      const doctorRegRequest = axios.post(registerDoctorURL, doctorData, { headers });
      const responseDataDoctor = await handleResponse(doctorRegRequest);

      console.log(responseDataDoctor.message);
      toast.success('Registration successfull')
      navigate(HOME_ENDPOINT);
    } catch (error) {
      console.error('Doctor Registration failed:', error.message);
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
        label="Last Name"
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
      <Select
        label="Specialization"
        name="specialization"
        value={formData.specialization}
        options={[
          // Adding an empty default option
          { label: 'Select Specialization', value: '' },
          ...ValidSpecializations.map((specialization) => ({
            value: specialization,
            label: specialization,
          })),
        ]}
        onChange={handleInputChange}
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

export default RegisterDoctorComponent;
