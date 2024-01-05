import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { GATEWAY_GET_PATIENT_BY_USER_ID, GATEWAY_GET_USER_BY_ID, UPDATE_PASSWORD_ENDPOINT } from "../../utils/endpoints";
import { buildURL, handleResponse } from "../../utils/utils";
import axios from "axios";
import { Spinner } from "react-bootstrap";

const PatientProfileComponent = ({userID, jwtToken}) => {
    const navigate = useNavigate();
    const [user, setUser] = useState();
    const [patientData, setPatientData] = useState();
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {  
        const headers = {
          Authorization: `Bearer ${jwtToken}`,
        };
    
        const fetchUser = async () => {
          try {
            const userByIDURL = buildURL(GATEWAY_GET_USER_BY_ID + userID, "");
            var request = axios.get(userByIDURL, {headers});
            var responseData = await handleResponse(request);
            setUser(responseData.payload);

            const patientByUserIDURL = buildURL(GATEWAY_GET_PATIENT_BY_USER_ID + userID, "")
            request = axios.get(patientByUserIDURL, {headers});
            responseData = await handleResponse(request);

            setPatientData(responseData.payload);
          } catch (error) {
            console.error('Error fetching user:', error.message);
            toast.error('Error getting profile..');
          } finally {
            setIsLoading(false);
          }
        };
    
        fetchUser();
      }, [navigate, jwtToken, userID]);

      const formatDate = (dateString) => {
        const options = { year: 'numeric', month: 'long', day: 'numeric' };
        return new Date(dateString).toLocaleDateString(undefined, options);
      };

      return (
        <div>
          {isLoading ? (
            // Loading state - render the spinner
            <Spinner animation="border" role="status" className="mx-auto"></Spinner>
          ) : (
            // Data loaded - render content
            <>
              <h1>Welcome, {user.username}!</h1>
              <p>
                If you'd like to change your password, click{" "}
                <span
                  style={{ color: "blue", cursor: "pointer" }}
                  onClick={() => navigate(UPDATE_PASSWORD_ENDPOINT)}
                >
                  here
                </span>
                .
              </p>
              <div>
                <h2>Patient Information:</h2>
                <p>First Name: {patientData.firstName}</p>
                <p>Last Name: {patientData.secondName}</p>
                <p>Email: {patientData.email}</p>
                <p>Phone Number: {patientData.phoneNumber}</p>
                <p>CNP: {patientData.cnp}</p>
                <p>BirthDay: {formatDate(patientData.birthDay)}</p>
              </div>
            </>
          )}
        </div>
      );    
}

export default PatientProfileComponent;
