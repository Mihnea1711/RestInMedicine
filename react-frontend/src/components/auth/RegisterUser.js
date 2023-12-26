import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import Input from '../common/Input'; // Reusable input component
import Select from '../common/Select'; // New Select component
import decodeSanitizedResponse from '../../services/Decoder';
import { roleOptions } from '../../utils/constants';
import { HOME_ENDPOINT, REGISTER_PATIENT_ENDPOINT, REGISTER_DOCTOR_ENDPOINT, GATEWAY_REGISTER_USER } from '../../utils/endpoints';
import { trimLastSegmentFromUrl } from '../../utils/utils';
import { validateRegistrationInput } from '../../validation/validateUserData';

const RegisterUser = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    role: ''
  });

  const [isSubmitting, setIsSubmitting] = useState(false);
  const navigate = useNavigate();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    // Add your registration logic here, e.g., API call or authentication
    console.log('Registration data:', formData);

    if (!validateRegistrationInput(formData)) {
      return;
    }

    setIsSubmitting(true);

    try {
      // Send POST request to the server
      const response = await axios.post(GATEWAY_REGISTER_USER, formData);
      // console.log('Registration successful:');
      const responseData = decodeSanitizedResponse(response.data);
      console.log(responseData);

      // Check if the response contains lastInsertedID
      if (true) {
        const id = 2;
        // Determine the redirect URL based on the selected role
        let redirectUrl;

        switch (formData.role) {
          case 'admin':
            redirectUrl = HOME_ENDPOINT;
            break;
          case 'patient':
            redirectUrl = trimLastSegmentFromUrl(REGISTER_PATIENT_ENDPOINT) + id;
            break;
          case 'doctor':
            redirectUrl = trimLastSegmentFromUrl(REGISTER_DOCTOR_ENDPOINT) + id;
            break;
          default:
            console.error('Invalid role:', formData.role);
            // Handle the case where an invalid role is selected
            return;
        }

        // Navigate to the determined redirect URL
        navigate(redirectUrl);
      } else {
        console.error('lastInsertedID not found in the response');
      }
      // Add any additional handling based on the server response
    } catch (error) {
      console.error('Registration failed:', error.message);
      // Add error handling, e.g., display an error message to the user
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
      <Select
        label="Role"
        name="role"
        value={formData.role}
        options={roleOptions}
        onChange={handleInputChange}
        required
      />
      <button type="submit" disabled={isSubmitting}>Register</button>
    </form>
  );
};

export default RegisterUser;
