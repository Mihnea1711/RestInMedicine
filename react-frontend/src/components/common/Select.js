import React from 'react';

const Select = ({ label, name, value, options, onChange, required }) => {
  return (
    <div>
      <label>
        {label}:
        <select name={name} value={value} onChange={onChange} required={required}>
          {options.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
      </label>
    </div>
  );
};

export default Select;
