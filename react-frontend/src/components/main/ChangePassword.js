import React, { useState } from 'react';
import Input from '../common/Input';
import { toast } from 'react-toastify';
import { buildURL, handleResponse } from '../../utils/utils';
import { GATEWAY_UPDATE_PASSWORD, PROFILE_ENDPOINT } from '../../utils/endpoints';
import axios from 'axios';
import { PasswordData } from '../../models/PasswordData';
import { useNavigate } from 'react-router-dom';

const ChangePasswordComponent = ({userID, jwt}) => {
  // State for form inputs
  const [newPassword, setNewPassword] = useState('');
  const [confirmNewPassword, setConfirmNewPassword] = useState('');
  const navigate = useNavigate();

  const headers = {
    Authorization: `Bearer ${jwt}`,
  };

  // Function to handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault();

    // Check if the new passwords match
    if (newPassword !== confirmNewPassword) {
      toast.error('New password does not match')
      return;
    }

    // Perform your password change logic here
    // (This is where you might make an API call to update the password)
    const changePasswordURL = buildURL(GATEWAY_UPDATE_PASSWORD(userID), "");
    const passwordData = new PasswordData(newPassword);

    try {
        const passwordRequest = axios.post(changePasswordURL, passwordData, {headers})
        const responseData = await handleResponse(passwordRequest);        
        toast.success(responseData.message);
        navigate(PROFILE_ENDPOINT);
        console.log("Password updated successfully");
    } catch (err) {
        console.error("An error occured while updating password." + err.message);
        toast.error(err.message);
    }

    // Reset form fields and error message on successful submission
    setNewPassword('');
    setConfirmNewPassword('');
  };

  return (
    <div className="container mt-4">
      <form onSubmit={handleSubmit}>
        <Input
          label="New Password"
          type="password"
          name="newPassword"
          value={newPassword}
          onChange={(e) => setNewPassword(e.target.value)}
          required
          className="form-control"
        />

        <Input
          label="Confirm New Password"
          type="password"
          name="confirmNewPassword"
          value={confirmNewPassword}
          onChange={(e) => setConfirmNewPassword(e.target.value)}
          required
          className="form-control"
        />

        <button type="submit" className="btn btn-primary mt-3">
          Change Password
        </button>
      </form>
    </div>
  );
};

export default ChangePasswordComponent;
