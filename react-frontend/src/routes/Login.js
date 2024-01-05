import React from 'react';
import LoginComponent from '../components/auth/Login';

const Login = () => {
  return (
    <div className="container mt-4">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card">
            <div className="card-body">
              <h2 className="card-title text-center mb-4">Login</h2>
              <LoginComponent />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
