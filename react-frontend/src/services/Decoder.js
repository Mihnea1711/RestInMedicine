import ResponseData from '../models/ResponseData';

function decodeSanitizedResponse(sanitizedResponse) {
    // Replace HTML entities with actual characters
    const decodedResponse = sanitizedResponse.replace(/&#34;/g, '"');
  
    // Parse the JSON string into an object
    const responseObject = JSON.parse(decodedResponse);
  
    // Extract values from the object and create a ResponseData instance
    const { message, error, payload, _links } = responseObject;
    const responseData = new ResponseData(message, error, payload, _links);
  
    return responseData;
}

export default decodeSanitizedResponse;