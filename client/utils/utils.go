package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func FromJSONToStruct[T any](resp *http.Response) (T, error) {
	var response T
	if resp.Body == nil {
		return response, fmt.Errorf("response body is empty")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return response, err
	}

	slog.Debug("response", slog.String("body", string(respBody)))

	return response, nil
}

func ToJSONBodyBuffer[T any](input T) (*bytes.Buffer, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	slog.Debug("request", slog.String("body", string(body)))

	return bytes.NewBuffer(body), nil
}

func ToStringBody(resp *http.Response) (string, error) {
	if resp.Body == nil {
		return "", fmt.Errorf("response body is empty")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
