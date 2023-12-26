import React, { useState } from 'react';
import axios from 'axios';
import Input from '../common/Input';
import Select from '../common/Select';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { validatePatientData } from '../../validation/validatePatientData';

const RegisterPatientComponent = () => {
  const [formData, setFormData] = useState({
    idUser: '',
    firstName: '',
    secondName: '',
    email: '',
    phoneNumber: '',
    cnp: '',
    birthDay: null,
    isActive: true,
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleDateChange = (date) => {
    setFormData({ ...formData, birthDay: date });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic form validation
    if (!validatePatientData(formData)) {
      return;
    }

    try {
      // Send POST request to the server
      const response = await axios.post('http://localhost:8080/api/patients', formData);
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
      <label>
        Birth Day:
        <DatePicker selected={formData.birthDay} onChange={handleDateChange} required />
      </label>
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

export default RegisterPatientComponent;
