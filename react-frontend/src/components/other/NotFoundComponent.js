import React from 'react';
import NotFoundImage from '../../static/404.jpg';

const NotFoundComponent = () => {
  return (
    <div className="text-center mt-5">
      <h2 className="text-danger">The requested page has not been found...</h2>
      <img src={NotFoundImage} alt="404 Not Found" className="img-fluid border border-secondary shadow p-3 mt-3" />
    </div>
  );
};

export default NotFoundComponent;
