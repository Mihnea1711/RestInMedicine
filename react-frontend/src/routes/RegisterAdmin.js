import React from 'react';
import RegisterAdminComponent from '../components/auth/RegisterAdmin';

const RegisterAdmin = () => {
  return (
    <div className="container mt-4">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card">
            <div className="card-body">
              <h2 className="card-title text-center mb-4">Enter Admin Data</h2>
              <RegisterAdminComponent />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegisterAdmin;
