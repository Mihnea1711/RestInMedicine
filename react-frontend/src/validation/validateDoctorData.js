import { toast } from 'react-toastify';
import { ValidSpecializations } from '../utils/constants';

export const validateDoctorData = (data) => {
  const { firstName, secondName, email, phoneNumber, specialization } = data;

  if (!firstName.trim()) {
    toast.error('First Name is required');
    return false;
  }

  if (!secondName.trim()) {
    toast.error('Second Name is required');
    return false;
  }

  if (!email.trim()) {
    toast.error('Email is required');
    return false;
  } else if (!isValidEmail(email)) {
    toast.error('Invalid email address');
    return false;
  }

  if (!phoneNumber.trim()) {
    toast.error('Phone Number is required');
    return false;
  } else if (!isValidPhoneNumber(phoneNumber)) {
    toast.error('Invalid phone number');
    return false;
  }

  if (!specialization) {
    toast.error('Specialization is required');
    return false;
  } else if (!ValidSpecializations.includes(specialization)) {
    toast.error('Invalid specialization');
    return false;
  }

  return true;
};

// Helper function to validate email address
const isValidEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

// Helper function to validate phone number
const isValidPhoneNumber = (phoneNumber) => {
  const phoneRegex = /^[0-9]{10}$/;
  return phoneRegex.test(phoneNumber);
};
