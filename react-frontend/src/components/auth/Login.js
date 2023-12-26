import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import { toast } from 'react-toastify';

import Input from '../common/Input'; // Update the path to the actual location of your Input component
import { GATEWAY_LOGIN, HOME_ENDPOINT } from '../../utils/endpoints';
import decodeSanitizedResponse from '../../services/Decoder';

const LoginComponent = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
  });

  const navigate = useNavigate(); // Hook for navigation

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic form validation
    if (!formData.username || !formData.password) {
      alert('Username and password are required');
      return;
    }

    try {
      console.log(formData);

      // Send POST request to the server for login
      const response = await axios.post(`${GATEWAY_LOGIN}`, formData);

      // Handle the login success, e.g., store user token, navigate to another page, etc.
      console.log('Login successful:', response.data);

      const responseData = decodeSanitizedResponse(response.data);
      console.log(responseData);

      // Check if the login was successful
      if (responseData.status === 200 && responseData.data.jwt) {
        // Store the JWT in cookies
        Cookies.set('jwt', response.data.jwt, { expires: 7 }); // Set the expiration as needed
        toast.success('Login successful');
        navigate(HOME_ENDPOINT);
      } else {
        toast.error('Login failed');
        console.error('Login failed:', response.data.message);
      }
    } catch (error) {
      console.error('Login failed:', error.message);
      // Handle login failure, e.g., display an error message to the user
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
      <button type="submit">Login</button>
    </form>
  );
};

export default LoginComponent;
