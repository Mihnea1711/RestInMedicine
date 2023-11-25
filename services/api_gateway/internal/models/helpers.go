package models

type RowsAffected struct {
	RowsAffected int `json:"rows_affected"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}
