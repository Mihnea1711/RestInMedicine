import { toast } from 'react-toastify';
import { roleOptions } from '../utils/constants';

export const validateRegistrationInput = (formData) => {
  const { username, password, role } = formData;

  if (!username || !password || !role) {
    toast.error('All fields are required');
    return false;
  }

  // Validate username
  const usernameRegex = /^[a-zA-Z0-9_-]{3,20}$/;
  if (!usernameRegex.test(username)) {
    toast.error('Invalid username format. It must be 3-20 characters and may contain letters, numbers, underscores, or hyphens.');
    return false;
  }

  // Validate password
  const passwordRegex = /^(?=.*\d)(?=.*[a-zA-Z]).{6,}$/;
  if (!passwordRegex.test(password)) {
    toast.error('Invalid password format. It must be at least 6 characters and include at least one letter and one number.');
    return false;
  }

  // Validate role
  const validRoles = roleOptions.map(option => option.value);
  if (!validRoles.includes(role)) {
    toast.error('Invalid role. Please select a valid role.');
    return false;
  }

  // Additional validation checks can be added as needed

  return true;
};
