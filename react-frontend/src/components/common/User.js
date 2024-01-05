import React from 'react';

export const User = ({ user, onDelete, onAddToBlacklist }) => {
  return (
    <tr key={user.idUser}>
      <td>{user.idUser}</td>
      <td>{user.username}</td>
      <td className='d-flex justify-content-center'>
        <button onClick={() => onDelete(user.idUser)} className="btn btn-danger btn-sm mx-2">Delete</button>
        {/* Uncomment the following line when you have the 'onAddToBlacklist' functionality */}
        {/* <button onClick={() => onAddToBlacklist(user.id)} className="btn btn-warning btn-sm">Add to Blacklist</button> */}
      </td>
    </tr>
  );
};
