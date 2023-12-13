package models

import (
	"net/http"
)

type RowsAffected struct {
	RowsAffected int `json:"rows_affected"`
}

type ResponseData struct {
	Message  string      `json:"message"`
	Error    string      `json:"error"`
	Payload  interface{} `json:"payload"`
	LinkList EndpointMap `json:"_links"`
}

type ResponseDataWrapper struct {
	Message string       `json:"message"`
	Error   string       `json:"error"`
	Payload interface{}  `json:"payload"`
	Header  *http.Header `json:"header"`
}

type EndpointData struct {
	Endpoint string `json:"href"`
	Method   string `json:"type"`
}

type LinkData struct {
	FieldName    string       `json:"fieldName"`
	EndpointData EndpointData `json:"data"`
}

// RedirectParams contains parameters for redirection.
type RedirectParams struct {
	Host     string
	Port     string
	Endpoint string
}

type EndpointMap map[string]EndpointData
