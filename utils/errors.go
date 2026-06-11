package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type APIError struct {
	StatusCode int
	Status     string `json:"status"`
	Errors     []struct {
		Code    string  `json:"code"`
		Message *string `json:"message,omitempty"`
	} `json:"errors"`
	RawBody []byte
}

func (e *APIError) Error() string {
	if len(e.Errors) > 0 {
		parts := make([]string, 0, len(e.Errors))
		for _, item := range e.Errors {
			msg := item.Code
			if item.Message != nil && *item.Message != "" {
				msg = item.Code + ": " + *item.Message
			}
			parts = append(parts, msg)
		}
		return strings.Join(parts, "; ")
	}
	if e.Status != "" {
		return e.Status
	}
	if len(e.RawBody) > 0 {
		return strings.TrimSpace(string(e.RawBody))
	}
	return "YANDEX MARKET PARTNER API ERROR"
}

func ParseAPIError(resp *http.Response) error {
	raw, _ := io.ReadAll(resp.Body)

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		RawBody:    raw,
	}

	_ = json.Unmarshal(raw, apiErr)

	if apiErr.Status == "" && len(apiErr.Errors) == 0 {
		apiErr.Status = strings.TrimSpace(string(raw))
	}

	return apiErr
}
