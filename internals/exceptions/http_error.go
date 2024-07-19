package exceptions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const responseWriterKey = 0

func NewError(ctx context.Context, errorId int) error {

	responseWriter := ctx.Value(responseWriterKey).(http.ResponseWriter)

	if responseWriter == nil {
		panic("Could not write on response context")
	}

	message := errorsMessages[errorId-1]

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

	return errors.New(message.message)
}

func NewErrorWithMessage(ctx context.Context, errorId int, messageParam string) error {

	responseWriter := ctx.Value(responseWriterKey).(http.ResponseWriter)

	if responseWriter == nil {
		panic("Could not write on response context")
	}

	message := errorsMessages[errorId-1]

	if !strings.Contains(message.message, "%s") {
		panic("The error handling don't have message parameter field")
	}

	message.message = fmt.Sprintf(message.message, messageParam)

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

	return errors.New(message.message)
}
