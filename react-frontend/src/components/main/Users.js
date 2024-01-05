import React, { useEffect, useMemo, useState } from 'react';
import axios from 'axios';
import { User } from '../common/User';
import { GATEWAY_GET_USERS, RABBIT_PUBLISH_ENDPOINT } from '../../utils/endpoints';
import Cookies from 'js-cookie';
import { JWT_COOKIE_NAME } from '../../utils/constants';
import { toast } from 'react-toastify';
import { buildRabbitURL, buildURL, handleResponse } from '../../utils/utils';
import { useNavigate } from 'react-router-dom';
import { Spinner, Table } from 'react-bootstrap';

const AllUsersComponent = ({claims, jwt}) => {
  const headers = useMemo(() => ({
    Authorization: `Bearer ${jwt}`,
  }), [jwt]);

  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const allUsersURL = buildURL(GATEWAY_GET_USERS, "");
        const request = axios.get(allUsersURL, { headers });
        const responseData = await handleResponse(request);

        setUsers(responseData.payload);
        setLoading(false);
        toast.success("Users loaded successfully.");
      } catch (error) {
        console.error('Error fetching users:', error.message);
        setLoading(false);
        toast.error('Error fetching users. Please try again.');
      }
    };

    fetchUsers();
  }, [jwt, navigate, headers]);

  const handleDelete = async (userId) => {
    const isCurrentUser = userId === parseInt(claims.sub, 10);

    // Show confirmation popup
    const confirmed = isCurrentUser
    ? window.confirm("Are you sure you want to delete yourself?")
    : window.confirm("Are you sure you want to delete this user?");

    if (confirmed) {
      const requestData = {
        "userID": userId
      }
      const deleteUserURL = buildRabbitURL(RABBIT_PUBLISH_ENDPOINT);
      try {
        const request = axios.post(deleteUserURL, requestData, { headers });
        await handleResponse(request);

        // Refresh the page after successful deletion
        if (isCurrentUser) {
          // Clear the "jwt" cookie
          Cookies.remove(JWT_COOKIE_NAME);
        }
        window.location.reload();
      } catch (err) {
        // Handle error
        console.error("Delete Error: ", err);
        toast.error("Error deleting user. Please try again.");
      }
    }
  };

  const handleAddToBlacklist = (userId) => {
    // Handle add to blacklist logic
    console.log(`Adding user with ID to blacklist: ${userId}`);
  };

  return (
    <div className="container mt-4">
       {loading ? (
        <Spinner animation="border" role="status" className="mx-auto"></Spinner>
      ) : (
        <Table striped bordered hover responsive className='text-center'>
          <thead>
            <tr>
              <th>ID</th>
              <th>Username</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <User
                key={user.idUser}
                user={user}
                onDelete={handleDelete}
                onAddToBlacklist={handleAddToBlacklist}
              />
            ))}
          </tbody>
        </Table>
        )
      }
    </div>
  );
};

export default AllUsersComponent;
