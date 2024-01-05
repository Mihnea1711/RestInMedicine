import React from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

const MyDatePicker = ({ disabled, selected, required, onChange }) => {
  return (
    <div className="form-group mt-2">
      <label htmlFor="birthDay">Birth Day:</label>
      <DatePicker
        id="birthDay"
        className="form-control"
        disabled={disabled}
        selected={selected}
        required={required}
        onChange={onChange}
      />
    </div>
  );
};

export default MyDatePicker;
