import React, { useState } from 'react';
import axios from 'axios';
import Input from '../common/Input';
import Select from '../common/Select';

import { validateDoctorData } from '../../validation/validateDoctorData';
import { specializations } from '../../utils/constants';

const RegisterDoctorComponent = () => {
  const [formData, setFormData] = useState({
    firstName: '',
    secondName: '',
    email: '',
    phoneNumber: '',
    specialization: '',
    isActive: true,
  });

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

    try {
      // Send POST request to the server
      const response = await axios.post('http://localhost:8080/api/doctors', formData);
      console.log('Registration successful:', response.data);
      // Add any additional handling based on the server response
    } catch (error) {
      console.error('Registration failed:', error.message);
      // Add error handling, e.g., display an error message to the user
    }
  };

  return (
    <form onSubmit={handleSubmit}>
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
        options={specializations}
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
      <button type="submit">Register</button>
    </form>
  );
};

export default RegisterDoctorComponent;
