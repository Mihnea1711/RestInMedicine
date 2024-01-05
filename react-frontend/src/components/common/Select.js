import React from 'react';

const Select = ({ label, name, value, options, onChange, required }) => {
  return (
    <div className="form-group">
      <label htmlFor={name}>{label}:</label>
      <select
        className="form-control"
        id={name}
        name={name}
        value={value}
        onChange={onChange}
        required={required}
      >
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  );
};

export default Select;
