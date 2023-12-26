// ANGULAR ENDPOINTS
export const HOME_ENDPOINT = "/home"
export const REGISTER_ENDPOINT = "/register";
export const LOGIN_ENDPOINT = "/login";

export const REGISTER_PATIENT_ENDPOINT = "/register-patient/:userID"
export const REGISTER_DOCTOR_ENDPOINT = "/register-doctor/:userID"
// Add more endpoints as needed

// GATEWAY ENDPOINTS
export const GATEWAY_REGISTER_USER = "http://localhost:8080/api/users";
export const GATEWAY_REGISTER_PATIENT = "http://localhost:8080/api/patients";
export const GATEWAY_REGISTER_DOCTOR = "http://localhost:8080/api/doctors";
export const GATEWAY_LOGIN = "http://localhost:8080/api/login";