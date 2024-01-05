// Investigation.js

import React from 'react';
import Input from '../common/Input';

const Investigation = ({ investigation, onChange }) => {
  const handleInvestigationChange = (event) => {
    const { name, value } = event.target;
    const processedValue = name === 'processingTime' ? parseInt(value, 10) : value;
    onChange(name, processedValue);
  };

  return (
    <div>
      <Input
        label="Name"
        type="text"
        name="name"
        value={investigation.name}
        onChange={handleInvestigationChange}
        required
        className="form-control mb-2"
      />

      <Input
        label="Processing Time"
        type="number"
        name="processingTime"
        value={investigation.processingTime}
        onChange={handleInvestigationChange}
        required
        className="form-control mb-2"
      />

      <Input
        label="Result"
        type="text"
        name="result"
        value={investigation.result}
        onChange={handleInvestigationChange}
        required
        className="form-control mb-2"
      />
    </div>
  )
};

export default Investigation;
