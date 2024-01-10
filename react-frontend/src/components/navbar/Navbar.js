import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import { GATEWAY_BLACKLIST_TOKEN, HOME_ENDPOINT, LOGIN_ENDPOINT } from '../../utils/endpoints';
import { JWT_COOKIE_NAME } from '../../utils/constants';
import { BlacklistData } from '../../models/BlacklistData';
import { buildURL, handleResponse } from '../../utils/utils';
import axios from 'axios';

export const Navbar = () => {
  const navigate = useNavigate();
  const jwtToken = Cookies.get(JWT_COOKIE_NAME);

  const handleLogout = async () => {
    // Implement logout logic here (e.g., clear JWT from cookies)
    const jwt = Cookies.get(JWT_COOKIE_NAME);
    const headers = {
      Authorization: `Bearer ${jwt}`,
    };
    // send request to add jwt to blacklist
    const blacklistRequest = new BlacklistData(jwt);
    // If editing an existing appointment, use the update endpoint
    const blacklistTokenURL = buildURL(GATEWAY_BLACKLIST_TOKEN, "");
    const request = axios.post(blacklistTokenURL, blacklistRequest, { headers });
    const responseData = await handleResponse(request);
    console.log(responseData.message);

    Cookies.remove(JWT_COOKIE_NAME);
    // Redirect to home after logout
    navigate(HOME_ENDPOINT);
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
      <div className="container">
        <Link className="navbar-brand" to={HOME_ENDPOINT}>
          NoPainNoGain Med
        </Link>
        <div className="navbar-collapse justify-content-end">
          <ul className="navbar-nav">
            {!jwtToken && (
              <>
                <li className="nav-item">
                  <Link className="nav-link" to={LOGIN_ENDPOINT}>
                    Login
                  </Link>
                </li>
              </>
            )}
            {jwtToken && (
              <li className="nav-item">
                <button className="nav-link btn btn-link" onClick={handleLogout}>
                  Logout
                </button>
              </li>
            )}
          </ul>
        </div>
      </div>
    </nav>
  );
};

// export default Navbar;
