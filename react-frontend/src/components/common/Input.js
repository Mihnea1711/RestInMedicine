import React from 'react';

const Input = ({ label, type, name, value, onChange, required }) => {
  return (
    <div>
      <label>
        {label}:
        <input
          type={type}
          name={name}
          value={value}
          onChange={onChange}
          required={required}
        />
      </label>
    </div>
  );
};

export default Input;
