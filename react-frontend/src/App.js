import React from 'react';
import {Routes, Route, Navigate} from "react-router-dom";

import Home from './routes/Home'
import NotFound from './routes/NotFound';
import { HOME_ENDPOINT, LOGIN_ENDPOINT, REGISTER_ADMIN_ENDPOINT, REGISTER_DOCTOR_ENDPOINT, REGISTER_PATIENT_ENDPOINT, PROFILE_ENDPOINT, UPDATE_PASSWORD_ENDPOINT, USERS_ENDPOINT, DOCTORS_ENDPOINT, APPOINTMENT_HISTORY_ENDPOINT, CONSULTATION_HISTORY_ENDPOINT, CREATE_APPOINTMENT_ENDPOINT, UPDATE_APPOINTMENT_URL, CREATE_CONSULTATION_ENDPOINT, UPDATE_CONSULTATION_URL } from './utils/endpoints';
import RegisterPatient from './routes/RegisterPatient';
import RegisterDoctor from './routes/RegisterDoctor';
import Login from './routes/Login';
import RegisterAdmin from './routes/RegisterAdmin';
import AllUsers from './routes/AllUsers';
import Profile from './routes/Profile';
import ChangePassword from './routes/ChangePassword';
import AllDoctors from './routes/AllDoctors';
import AppointmentHistory from './routes/AppointmentHistory';
import ConsultationHistory from './routes/ConsultationHistory';
import CreateEditAppointment from './routes/CreateEditAppointment';
import CreateEditConsultation from './routes/CreateEditConsultation';
import { Navbar } from './components/navbar/Navbar';

const App = () => {
  return (
    <div>
      <Navbar/>
      <Routes>
        <Route path="/" element={<Navigate to={HOME_ENDPOINT} />} />

        <Route path={HOME_ENDPOINT} element={<Home />} />
        <Route path={LOGIN_ENDPOINT} element={<Login />} />
        <Route path={REGISTER_ADMIN_ENDPOINT} element={<RegisterAdmin />} />
        <Route path={REGISTER_PATIENT_ENDPOINT} element={<RegisterPatient />} />
        <Route path={REGISTER_DOCTOR_ENDPOINT} element={<RegisterDoctor />} />

        <Route path={PROFILE_ENDPOINT} element={<Profile />} />
        <Route path={UPDATE_PASSWORD_ENDPOINT} element={<ChangePassword />} />

        <Route path={USERS_ENDPOINT} element={<AllUsers />} />
        <Route path={DOCTORS_ENDPOINT} element={<AllDoctors />} />    
        <Route path={APPOINTMENT_HISTORY_ENDPOINT} element={<AppointmentHistory />} />   
        <Route path={CONSULTATION_HISTORY_ENDPOINT} element={<ConsultationHistory />} /> 

        <Route path={CREATE_APPOINTMENT_ENDPOINT} element={<CreateEditAppointment />} />    
        <Route path={UPDATE_APPOINTMENT_URL} element={<CreateEditAppointment />} />    

        <Route path={CREATE_CONSULTATION_ENDPOINT} element={<CreateEditConsultation />} />    
        <Route path={UPDATE_CONSULTATION_URL} element={<CreateEditConsultation />} />  

        <Route path="*" element={<NotFound />} />
      </Routes>
    </div>
  );
};

export default App;
