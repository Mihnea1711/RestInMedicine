import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import { toast } from 'react-toastify';

import Input from '../common/Input'; // Update the path to the actual location of your Input component
import { GATEWAY_LOGIN, HOME_ENDPOINT } from '../../utils/endpoints';
import { JWT_COOKIE_DURATION_DAYS, JWT_COOKIE_NAME, JWT_COOKIE_SAME_SITE, JWT_COOKIE_SECURE_FLAG } from '../../utils/constants';
import { LoginData } from '../../models/LoginData';
import { buildURL, handleResponse } from '../../utils/utils';
import decodeSanitizedResponse from '../../services/Decoder';

const LoginComponent = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
  });

  const navigate = useNavigate();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic form validation
    if (!formData.username || !formData.password) {
      toast.error('Username and password are required');
      return;
    }

    try {
      const loginData = new LoginData(formData.username, formData.password); 

      // Send POST request to the server for login
      const loginURL = buildURL(GATEWAY_LOGIN, "");
      const responseData = await handleResponse(axios.post(loginURL, loginData));

      const jwt = responseData.payload;
      // Store the JWT in cookies
      Cookies.set(JWT_COOKIE_NAME, jwt, { expires: JWT_COOKIE_DURATION_DAYS, sameSite: JWT_COOKIE_SAME_SITE, secure: JWT_COOKIE_SECURE_FLAG });
      toast.success('Login successful');
      navigate(HOME_ENDPOINT);
      
    } catch (error) {
      console.error('Login failed:', error.message);
      toast.error(decodeSanitizedResponse(error.response.data).message);
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
      <button type="submit" className="btn btn-primary btn-block mt-4">Login</button>
    </form>
  );
};

export default LoginComponent;
