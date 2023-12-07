package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

type ResponseInterceptor struct {
	http.ResponseWriter
	request *http.Request
}

func (ri *ResponseInterceptor) Write(data []byte) (int, error) {
	var responseData models.ResponseData
	err := json.Unmarshal(data, &responseData)
	if err == nil {
		responseData.LinkList = utils.GetHateoasData(ri.request.URL.Path, ri.request.Method)
		updatedData, err := json.Marshal(responseData)
		if err == nil {
			return ri.ResponseWriter.Write(updatedData)
		} else {
			log.Printf("Error marshaling updated response data: %v", err)
		}
	} else {
		log.Printf("Error unmarshaling response data: %v", err)
	}

	return ri.ResponseWriter.Write(data)
}

// Middleware function to intercept response and add path and method
func AddPathAndMethodToResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		interceptor := &ResponseInterceptor{ResponseWriter: w, request: r}
		next.ServeHTTP(interceptor, r)
	})
}
