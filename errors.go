package main

import "fmt"

type ClientError struct {
	Endpoint        string
	StatusCode      int
	ResponseContent []byte
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("Request to %s unsuccessful. Status code: %d. Response content: %s", c.Endpoint, c.StatusCode, c.ResponseContent)
}

func NewClientError(endpoint string, statusCode int, responseContent []byte) *ClientError {
	return &ClientError{
		Endpoint:        endpoint,
		StatusCode:      statusCode,
		ResponseContent: responseContent,
	}
}

// InvalidParamsError message: str, key: str, value: Any
type InvalidParamsError struct {
	Message string
	key     string
	value   any
}

func NewInvalidParamsError(message, key string, value any) *InvalidParamsError {
	return &InvalidParamsError{
		Message: message,
		key:     key,
		value:   value,
	}
}

func (c *InvalidParamsError) Error() string {
	return fmt.Sprintf("Invalid params for %s: %s. Message: %s", c.key, c.value, c.Message)
}
