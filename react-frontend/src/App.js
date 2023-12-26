import React from 'react';
import {Routes, Route, Navigate} from "react-router-dom";

import Registration from './routes/Register';
import Home from './routes/Home'
import NotFound from './routes/NotFound';
import { HOME_ENDPOINT, LOGIN_ENDPOINT, REGISTER_DOCTOR_ENDPOINT, REGISTER_ENDPOINT, REGISTER_PATIENT_ENDPOINT } from './utils/endpoints';
import RegisterPatient from './routes/RegisterPatient';
import RegisterDoctor from './routes/RegisterDoctor';
import Login from './routes/Login';

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/home" />} />

      <Route path={HOME_ENDPOINT} element={<Home />} />
      <Route path={REGISTER_ENDPOINT} element={<Registration />} />
      <Route path={REGISTER_PATIENT_ENDPOINT} element={<RegisterPatient />} />
      <Route path={REGISTER_DOCTOR_ENDPOINT} element={<RegisterDoctor />} />
      <Route path={LOGIN_ENDPOINT} element={<Login />} />


      <Route path="*" element={<NotFound />} />
    </Routes>
  );
};

export default App;
