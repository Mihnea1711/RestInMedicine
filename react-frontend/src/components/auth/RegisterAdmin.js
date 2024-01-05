import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import axios from 'axios';
import Input from '../common/Input';
import { RegisterData } from '../../models/RegisterData';
import { validateRegistrationInput } from '../../validation/validateUserData';
import { buildURL, handleResponse } from '../../utils/utils';
import { ROLE_ADMIN } from '../../utils/constants';
import { GATEWAY_REGISTER_USER, HOME_ENDPOINT } from '../../utils/endpoints';

const RegisterAdminComponent = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    role: ROLE_ADMIN
  });

  const [isSubmitting, setIsSubmitting] = useState(false);
  const navigate = useNavigate();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!validateRegistrationInput(formData)) {
      return;
    }

    setIsSubmitting(true);

    try {
      var adminUserReg = new RegisterData(formData.username, formData.password, formData.role);

      const registerUserURL = buildURL(GATEWAY_REGISTER_USER, "");
      const userRegRequest = axios.post(registerUserURL, adminUserReg);
      const responseDataUser = await handleResponse(userRegRequest);

      if (responseDataUser.payload) {
        toast.success('Registration successful');
        navigate(HOME_ENDPOINT);
      } else {
        console.error('Admin Registration failed:', responseDataUser.message);
        toast.error(`Registration failed: ${responseDataUser.message}`);
      }
    } catch (error) {
      console.error('Doctor Registration failed:', error.message);
      toast.error(`Error during registration: ${error.message}`);
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
      <button type="submit" disabled={isSubmitting} className="btn btn-primary btn-block mt-4">Register</button>
    </form>
  );
};

export default RegisterAdminComponent;
