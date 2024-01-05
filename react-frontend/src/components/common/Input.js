import React, { useState } from 'react';

const Input = ({ label, type, name, value, onChange, required, disabled, placeholder }) => {
  const [showPassword, setShowPassword] = useState(false);

  const togglePasswordVisibility = () => {
    setShowPassword((prevShowPassword) => !prevShowPassword);
  };

  return (
    <div className="form-group">
      <label htmlFor={name}>{label}</label>
      {type === 'password' ? (
        <div className="input-group">
          <input
            type={showPassword ? 'text' : 'password'}
            id={name}
            name={name}
            value={value}
            onChange={onChange}
            required={required}
            disabled={disabled}
            className="form-control"
            placeholder={placeholder}
          />
          <div className="input-group-append">
            <button
              className="btn btn-outline-secondary"
              type="button"
              onClick={togglePasswordVisibility}
            >
              {showPassword ? 'Hide' : 'Show'}
            </button>
          </div>
        </div>
      ) : (
        <input
          type={type}
          id={name}
          name={name}
          value={value}
          onChange={onChange}
          required={required}
          disabled={disabled}
          className="form-control"
          placeholder={placeholder}
        />
      )}
    </div>
  );
};

export default Input;
