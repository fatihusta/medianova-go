package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/fatihusta/medianova-go/common"
)

func DoHTTPRequest[T any](c *http.Client, req *http.Request) *common.Result[T] {

	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		errTempl := fmt.Sprintf("method:%s, url:%s",
			req.Method, req.URL.Scheme+"://"+req.URL.Host+req.URL.Path)
		result := common.NewResult[T]()
		result.Error = fmt.Errorf("%s, %s", err.Error(), errTempl)
		return result
	}
	defer resp.Body.Close()

	return Result[T](resp)
}

func Result[T any](resp *http.Response) *common.Result[T] {
	result := common.NewResult[T]()

	result.Status = resp.StatusCode

	errTempl := fmt.Sprintf("method:%s, url:%s",
		resp.Request.Method, resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+resp.Request.URL.Path)

	if resp.StatusCode != http.StatusOK {
		errMsg, err := ToStringBody(resp)
		if err == nil {
			result.Error = fmt.Errorf("request not succeeded, error:%s, %s", errMsg, errTempl)
		} else {
			result.Error = fmt.Errorf("request not succeeded, %s", errTempl)
		}

		return result
	}

	if resp.Body == nil {
		result.Error = fmt.Errorf("response body is empty, %s", errTempl)
		return result
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Error = fmt.Errorf("%s, %s", err.Error(), errTempl)
		return result
	}

	err = json.Unmarshal(respBody, &result.Body)
	if err != nil {
		result.Error = fmt.Errorf("%s, %s, body:%s", err.Error(), errTempl, respBody)
		return result
	}

	result.Headers = resp.Header

	slog.Debug("result", slog.String("body", string(respBody)))

	return result
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

// Decode to struct
func DecodeToStruct[T any](input any) (T, error) {
	var result T

	// to json
	b, err := json.Marshal(input)
	if err != nil {
		return result, fmt.Errorf("marshal failed: %w", err)
	}

	// to struct
	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, fmt.Errorf("unmarshal failed: %w", err)
	}

	return result, nil
}
