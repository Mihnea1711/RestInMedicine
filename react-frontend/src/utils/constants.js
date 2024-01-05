export const roleOptions = [
    { label: 'Select Role', value: '' },
    { label: 'Admin', value: 'admin' },
    { label: 'Patient', value: 'patient' },
    { label: 'Doctor', value: 'doctor' },
];

const StatusAppointment = {
  SCHEDULED: 'scheduled',
  CONFIRMED: 'confirmed',
  NOT_PRESENT: 'not_present',
  CANCELED: 'canceled',
  HONORED: 'honored',
};

// ValidStatus array
export const ValidStatus = Object.values(StatusAppointment);

const Specialization = {
  CARDIOLOGY: 'Cardiology',
  NEUROLOGY: 'Neurology',
  ORTHOPEDICS: 'Orthopedics',
  PEDIATRICS: 'Pediatrics',
  DERMATOLOGY: 'Dermatology',
  RADIOLOGY: 'Radiology',
  SURGERY: 'Surgery',
};

// ValidSpecializations array
export const ValidSpecializations = Object.values(Specialization);

export const JWT_COOKIE_NAME = "jwt";
export const JWT_COOKIE_DURATION_DAYS = 1;
export const JWT_COOKIE_SAME_SITE = "None";
export const JWT_COOKIE_SECURE_FLAG = true;
export const JWT_SECRET = "thisshouldbeabettersecret";

export const GATEWAY_SCHEME = "http://";
export const GATEWAY_HOST = "localhost";
export const GATEWAY_PORT = 8080;

export const RABBIT_SCHEME = "http://"
export const RABBIT_HOST = "localhost"
export const RABBIT_PORT = 8090

export const ROLE_ADMIN = "admin";
export const ROLE_DOCTOR = "doctor";
export const ROLE_PATIENT = "patient";

export const TIMESTAMP_SUFFIX = "T00:00:00Z"

export const QUERY_DOCTOR_ID = "doctorID";
export const QUERY_PATIENT_ID = "patientID";
