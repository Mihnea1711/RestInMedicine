class ResponseData {
    constructor(message, error, payload, links) {
        this.message = message;
        this.error = error;
        this.payload = payload;
        this.links = links;
    }
}

export default ResponseData;