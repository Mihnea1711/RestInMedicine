import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { GATEWAY_GET_USER_BY_ID, UPDATE_PASSWORD_ENDPOINT } from "../../utils/endpoints";
import { buildURL, handleResponse } from "../../utils/utils";
import axios from "axios";

const AdminProfileComponent = ({userID, jwtToken}) => {
    const navigate = useNavigate();
    const [user, setUser] = useState();

    useEffect(() => {    
      const headers = {
        Authorization: `Bearer ${jwtToken}`,
      };
  
      const fetchUser = async () => {
        try {
          const userByIDURL = buildURL(GATEWAY_GET_USER_BY_ID + userID, "");
          const request = axios.get(userByIDURL, {headers});
          const responseData = await handleResponse(request);
  
          setUser(responseData.payload);
        } catch (error) {
          console.error('Error fetching user:', error.message);
          toast.error('Error getting profile..');
        }
      };
    
      fetchUser();
    }, [navigate, jwtToken, userID]);

    return (
      <div className="container mt-4">
        {user && (
          <>
            <h1 className="mb-3">Welcome, {user.username}!</h1>
            <p>
              If you'd like to change your password, click{" "}
              <span
                className="text-primary"
                style={{ cursor: "pointer" }}
                onClick={() => navigate(UPDATE_PASSWORD_ENDPOINT)}
              >
                here
              </span>
              .
            </p>
          </>
        )}
      </div>
    );    
}

export default AdminProfileComponent;
