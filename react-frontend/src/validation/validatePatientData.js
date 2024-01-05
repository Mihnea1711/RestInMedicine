import { toast } from 'react-toastify';

// Validate patient data
export const validatePatientData = (data) => {
    const {username ,password, firstName, secondName, email, phoneNumber, cnp, birthDay} = data;

    if (!username || !password || !firstName || !secondName || !email || !phoneNumber || !cnp || !birthDay) {
        toast.error('All fields are required');
        return false;
      }
  
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
    } else if (!isValidPhoneNumber(phoneNumber)) {
        toast.error('Invalid phone number');
        return false;
    }
  
    if (!cnp.trim()) {
        toast.error('CNP is required');
        return false;
    } else if (!isValidCNP(cnp)) {
        toast.error('Invalid CNP');
        return false;
    }
  
    if (!birthDay) {
        toast.error('Birth Day is required');
        return false;
    }
  
    return true;
  };
  
  // Helper function to validate email address
  const isValidEmail = (email) => {
    // You can implement a more sophisticated email validation if needed
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };
  
  // Helper function to validate phone number
  const isValidPhoneNumber = (phoneNumber) => {
    // You can implement a more sophisticated phone number validation if needed
    const phoneRegex = /^[0-9]{10}$/;
    return phoneRegex.test(phoneNumber);
  };
  
  // Helper function to validate CNP (personal identification number)
  export const isValidCNP = (cnp) => {
    // You can implement a more sophisticated CNP validation if needed
    const cnpRegex = /^[0-9]{13}$/;
    return cnpRegex.test(cnp);
  };