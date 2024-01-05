import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import { HOME_ENDPOINT, LOGIN_ENDPOINT } from '../../utils/endpoints';
import { JWT_COOKIE_NAME } from '../../utils/constants';

export const Navbar = () => {
  const navigate = useNavigate();
  const jwtToken = Cookies.get(JWT_COOKIE_NAME);

  const handleLogout = () => {
    // Implement logout logic here (e.g., clear JWT from cookies)
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
