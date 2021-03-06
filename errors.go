package hgApi

import "fmt"

type APIError struct {
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("hosted graphite API : %v", e.Message)
}

func handleError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}

	if apiError.Message != "" {
		return apiError
	}

	return nil
}
