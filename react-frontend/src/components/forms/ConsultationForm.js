import React, { useEffect, useState } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import Input from '../common/Input';
import Textarea from '../common/Textarea';
import { ConsultationData } from '../../models/Consultation';
import Investigation from '../common/Investigation';

const ConsultationForm = ({ onSubmit, initialData }) => {
  const [formData, setFormData] = useState(initialData || new ConsultationData('', '', '', '', '', []));

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

  const handleInvestigationChange = (index, field, value) => {
    setFormData((prevData) => {
      const investigations = [...prevData.investigations];
      investigations[index][field] = value;
      return {
        ...prevData,
        investigations,
      };
    });
  };

  const handleRemoveInvestigation = (index) => {
    setFormData((prevData) => {
      const investigations = [...prevData.investigations];
      investigations.splice(index, 1);
      return {
        ...prevData,
        investigations,
      };
    });
  };

  const handleAddInvestigation = () => {
    setFormData((prevData) => ({
      ...prevData,
      investigations: [...prevData.investigations, { name: '', processingTime: 0, result: '' }],
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
    <form onSubmit={handleSubmit} className="mt-4">
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
        <DatePicker selected={formData.date ? new Date(formatDate(formData.date)) : null} onChange={handleDateChange} minDate={new Date()} required className="form-control" />
      </div>

      <Textarea
        label="Diagnostic"
        name="diagnostic"
        value={formData.diagnostic}
        onChange={handleChange}
        required
      />

      <div className="mb-3">
        <label className="mr-2">Investigations:</label>
        {formData.investigations.map((investigation, index) => (
          <div key={index} className="mb-3">
            <Investigation
              investigation={investigation}
              onChange={(field, value) => handleInvestigationChange(index, field, value)}
            />
            <button type="button" onClick={() => handleRemoveInvestigation(index)} className="btn btn-danger mt-2">
              Remove Investigation
            </button>
          </div>
        ))}
        <button type="button" onClick={handleAddInvestigation} className="btn btn-primary">
          Add Investigation
        </button>
      </div>

      <button type="submit" className="btn btn-success">Submit</button>
    </form>
  );
};

export default ConsultationForm;
