// Textarea.js

import React from 'react';

const Textarea = ({ label, name, value, onChange, required, className }) => {
  return (
    <div className={className}>
      <label>
        {label}:
        <textarea
          name={name}
          value={value}
          onChange={onChange}
          required={required}
          className="form-control"
        />
      </label>
    </div>
  );
};

export default Textarea;
