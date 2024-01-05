import decodeSanitizedResponse from "../services/Decoder";
import { GATEWAY_HOST, GATEWAY_PORT, GATEWAY_SCHEME, RABBIT_HOST, RABBIT_PORT, RABBIT_SCHEME } from "./constants";

export function parseJwt (token) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));
  
    return JSON.parse(jsonPayload);
}

export function buildURL(endpoint, query) {
    return GATEWAY_SCHEME + GATEWAY_HOST + ":" + GATEWAY_PORT + endpoint + query;
}

export function buildQueryString(params) {
    const queryString = Object.keys(params)
      .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`)
      .join('&');
  
    return `?${queryString}`;
}

export function buildRabbitURL(rabbit_endpoint) {
  return RABBIT_SCHEME + RABBIT_HOST + ":" + RABBIT_PORT + rabbit_endpoint;
}

export const handleResponse = async (requestPromise) => {
    try {
      const response = await requestPromise;
      const decodedResponse = decodeSanitizedResponse(response.data);
  
      if (response.status >= 200 && response.status < 300) {
        // Successful response
        return decodedResponse;
      }
      console.log("im herere");
      // Handle specific error status codes
      if (response.status === 400) {
        console.error(`Bad Request: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Bad Request: ${decodedResponse.message}`);
      } else if (response.status === 401) {
        console.error(`Unauthorized: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Unauthorized: ${decodedResponse.message}`);
      } else if (response.status === 403) {
        console.error(`Forbidden: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Forbidden: ${decodedResponse.message}`);
      } else if (response.status === 404) {
        console.error(`Not Found: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Not Found: ${decodedResponse.message}`);
      } else if (response.status === 422) {
        console.error(`Unprocessable Entity: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Unprocessable Entity: ${decodedResponse.message}`);
      } else if (response.status === 500) {
        console.error(`Internal Server Error: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Internal Server Error: ${decodedResponse.message}`);
      } else if (response.status === 502) {
        console.error(`Bad Gateway: ${decodedResponse.message}`, decodedResponse.error);
        throw new Error(`Bad Gateway: ${decodedResponse.message}`);
      }
  
      // If none of the above conditions are met, throw a generic error
      throw new Error(`Unhandled status code: ${response.status}`);
    } catch (error) {
      console.error('Request failed:', error.message);
      throw error;
    }
};

export function generateBirthdayFromCNP(cnp) {
  // Extract the birthdate information from the CNP
  const year = parseInt(cnp.substring(1, 3), 10);
  const month = parseInt(cnp.substring(3, 5), 10);
  const day = parseInt(cnp.substring(5, 7), 10);
  
  // Determine the century based on the first digit of the CNP
  const firstDigit = parseInt(cnp.charAt(0), 10);
  let century;
  if (firstDigit === 1 || firstDigit === 2) {
    century = 20;
  } else if (firstDigit === 3 || firstDigit === 4) {
    century = 19;
  } else if (firstDigit === 5 || firstDigit === 6) {
    century = 21;
  }
  
  // Create a Date object using the extracted information
  const birthday = new Date((century-1) * 100 + year, month - 1, day);
  
  return birthday;
}

export function verifyJWTRole(claims, allowedRoles) {
  // Check if the 'role' claim is present in the JWT payload
  if (!claims.role) {
    console.error('JWT does not contain a role claim.');
    return false;
  }

  // Check if the JWT role is in the list of allowed roles
  return allowedRoles.includes(claims.role);
}

export function getSubFromJWT(token) {
  try {
    const claims = parseJwt(token);

    // Check if the 'sub' claim is present in the JWT payload
    if (!claims.sub) {
      console.error('JWT does not contain a subject claim.');
      return false;
    }

    return claims.sub;
  } catch (error) {
    // Handle the case when there's an error decoding the JWT
    console.error('Error decoding JWT:', error.message);
    return null;
  }
}

export const validateJwtToken = (token) => {
  try {
    const decodedClaims = parseJwt(token);

    if (!decodedClaims.sub) {
      throw new Error('Invalid sub in JWT token');
    }

     // Check if 'sub' (user ID) is a valid integer
     const userId = parseInt(decodedClaims.sub, 10); // Parse 'sub' as an integer
     if (!Number.isInteger(userId)) {
       console.error('Invalid user ID in JWT token');
       throw new Error('Invalid user ID in JWT token');
     }

    // Check if 'exp' (expiration date) is a valid date in the future
    const currentTimestamp = Math.floor(Date.now() / 1000);
    if (decodedClaims.exp && decodedClaims.exp < currentTimestamp) {
      throw new Error('JWT token has expired');
    }

    // Validation passed, return the claims
    return decodedClaims;
  } catch (error) {
    throw new Error(`Error decoding/validating JWT token: ${error.message}`);
  }
};

