package models

type RowsAffected struct {
	RowsAffected int `json:"rows_affected"`
}

type ResponseData struct {
	Message  string      `json:"message"`
	Error    string      `json:"error"`
	Payload  interface{} `json:"payload"`
	LinkList []LinkData  `json:"_links"`
}

type EndpointData struct {
	Endpoint string `json:"href"`
	Method   string `json:"type"`
}

type LinkData struct {
	FieldName    string       `json:"field_name"`
	EndpointData EndpointData `json:"data"`
}
