import React, { useEffect, useState } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import Input from '../common/Input';
import Select from '../common/Select';
import { ValidStatus } from '../../utils/constants';

const AppointmentForm = ({ onSubmit, initialData }) => {
  const [formData, setFormData] = useState(initialData || { idPatient: '', date: '', status: '' });

  useEffect(() => {
    if (initialData) {
      setFormData(initialData);
    }
  }, [initialData]);

  const handleChange = (event) => {
    const { name, value } = event.target;

    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleDateChange = (date) => {
    setFormData((prevData) => ({
      ...prevData,
      date,
    }));
  };

  const handleSubmit = (event) => {
    event.preventDefault();

    onSubmit(formData);
  };

  const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString(undefined, options);
  };

  return (
    <form onSubmit={handleSubmit}>
      <Input
        label="ID Patient"
        type="text"
        name="idPatient"
        value={initialData ? initialData.idPatient : formData.idPatient}
        onChange={handleChange}
        disabled={initialData !== null}
        className="form-control mb-3"
      />

      <div className="mb-3">
        <label className="me-2 mt-3">Date:</label>
        <DatePicker className="form-control" selected={formData.date ? new Date(formatDate(formData.date)) : null} onChange={handleDateChange} minDate={new Date()} required />
      </div>

      <Select 
        label="Status"
        name="status"
        value={formData.status}
        onChange={handleChange}
        options={ValidStatus.map((status) => ({ value: status, label: status.toUpperCase() }))}
        required
      />

      <button type="submit" className="btn btn-success mt-3">Submit</button>
    </form>
  );
};

export default AppointmentForm;
