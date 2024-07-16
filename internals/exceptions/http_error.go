package exceptions

import (
	"context"
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const responseWriterKey = 0

func NewError(ctx context.Context, errorId int) {

	responseWriter := ctx.Value(responseWriterKey).(http.ResponseWriter)

	if responseWriter == nil {
		panic("Could not write on response context")
	}

	message := errors[errorId-1]

	errResponse := map[string]interface{}{
		"id":     message.id,
		"error":  message.message,
		"status": http.StatusBadRequest,
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(responseWriter).Encode(errResponse); err != nil {
		panic(err)
	}
	if flusher, ok := responseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}
