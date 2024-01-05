// ANGULAR ENDPOINTS
export const HOME_ENDPOINT = "/home"
export const LOGIN_ENDPOINT = "/login";
export const REGISTER_ADMIN_ENDPOINT = "/register-admin";

// ADMIN
export const REGISTER_PATIENT_ENDPOINT = "/register-patient";
export const REGISTER_DOCTOR_ENDPOINT = "/register-doctor";
export const USERS_ENDPOINT = "/users";
export const GET_USER_BY_ID_ENDPOINT = "/users/:userID";
export const UPDATE_USER_ENDPOINT = "/users/:userID/edit";

// PATIENT
export const APPOINTMENT_HISTORY_ENDPOINT = "/appointments";
export const CONSULTATION_HISTORY_ENDPOINT = "/consultations";
export const DOCTORS_ENDPOINT = "/doctors";

// DOCTOR
export const PATIENTS_ENDPOINTS = "/patients";
export const PATIENT_ENDPOINTS = "/patients/:patientID";

export const CREATE_APPOINTMENT_ENDPOINT = "/appointments/new";
export const GET_APPOINTMENT_ENDPOINT = "/appointments/:appointmentID";
export const UPDATE_APPOINTMENT_URL = "/appointments/:appointmentID/edit"
export const UPDATE_APPOINTMENT_ENDPOINT = (appointmentID) => `/appointments/${appointmentID}/edit`;

export const CREATE_CONSULTATION_ENDPOINT = "/consultations/new";
export const GET_CONSULTATION_ENDPOINT = "/consultations/:consultationID";
export const UPDATE_CONSULTATION_URL = "/consultations/:consultationID/edit"
export const UPDATE_CONSULTATION_ENDPOINT = (consultationID) => `/consultations/${consultationID}/edit`

// SHARED
export const PROFILE_ENDPOINT = "/profile";
export const UPDATE_PASSWORD_ENDPOINT = "/profile/update-password"

// GATEWAY ENDPOINTS
export const GATEWAY_REGISTER_USER = "/api/users";
export const GATEWAY_REGISTER_PATIENT = "/api/patients";
export const GATEWAY_REGISTER_DOCTOR = "/api/doctors";
export const GATEWAY_LOGIN = "/api/login";

export const GATEWAY_GET_USERS = "/api/users";
export const GATEWAY_GET_USER_BY_ID = "/api/users/";
export const GATEWAY_UPDATE_USER = "/api/users/";

export const GATEWAY_CREATE_APPOINTMENT = "/api/appointments";
export const GATEWAY_GET_APPOINTMENTS = "/api/appointments";
export const GATEWAY_GET_APPOINTMENT = "/api/appointments/";
export const GATEWAY_UPDATE_APPOINTMENT = "/api/appointments/";
export const GATEWAY_DELETE_APPOINTMENT = "/api/appointments/";

export const GATEWAY_CREATE_CONSULTATION = "/api/consultations";
export const GATEWAY_GET_CONSULTATIONS = "/api/consultations";
export const GATEWAY_GET_CONSULTATION = "/api/consultations/";
export const GATEWAY_UPDATE_CONSULTATION = "/api/consultations/";
export const GATEWAY_DELETE_CONSULTATION = "/api/consultations/";

export const GATEWAY_CREATE_DOCTOR = "/api/doctors";
export const GATEWAY_GET_DOCTORS = "/api/doctors";
export const GATEWAY_GET_DOCTOR = "/api/doctors/";
export const GATEWAY_GET_DOCTOR_BY_USER_ID = "/api/doctors/users/";
export const GATEWAY_UPDATE_DOCTOR = "/api/doctors/";

export const GATEWAY_CREATE_PATIENT = "/api/patients"
export const GATEWAY_GET_PATIENTS = "/api/patients"
export const GATEWAY_GET_PATIENT = "/api/patients/"
export const GATEWAY_GET_PATIENT_BY_USER_ID = "/api/patients/users/";
export const GATEWAY_UPDATE_PATIENT = "/api/patients/";

export const GATEWAY_UPDATE_PASSWORD = (userID) => `/api/users/${userID}/update-password`;

export const RABBIT_PUBLISH_ENDPOINT = "/api/rabbit/publish"
