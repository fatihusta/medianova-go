package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/fatihusta/medianova-go/common"
	"github.com/google/uuid"
)

const requestIDKey common.CtxKey = "request_id"

func DoHTTPRequest[T any](c *http.Client, req *http.Request) *common.Result[T] {

	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	reqID := uuid.New().String()
	req = req.WithContext(context.WithValue(req.Context(), requestIDKey, reqID))

	if slog.Default().Enabled(context.Background(), slog.LevelDebug) ||
		slog.Default().Enabled(context.Background(), slog.LevelInfo) {

		attrs := []slog.Attr{
			slog.String("request_id", reqID),
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
		}

		logLevel := slog.LevelInfo

		if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
			logLevel = slog.LevelDebug
			attrs = append(attrs, slog.String("body", ReqBodyToString(req)))
		}

		slog.LogAttrs(context.Background(), logLevel, "starting request", attrs...)
	}

	resp, err := c.Do(req)
	if err != nil {
		errTempl := fmt.Sprintf("%s: %s, method: %s, url: %s",
			requestIDKey,
			reqID,
			req.Method,
			req.URL.String())
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

	reqID := GetRequestID(resp.Request.Context())
	errTempl := fmt.Sprintf("%s: %s, method: %s, url: %s",
		requestIDKey,
		reqID,
		resp.Request.Method,
		resp.Request.URL.String())

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

	if slog.Default().Enabled(context.Background(), slog.LevelDebug) ||
		slog.Default().Enabled(context.Background(), slog.LevelInfo) {

		attrs := []slog.Attr{
			slog.String(requestIDKey.String(), reqID),
			slog.Int("status", result.Status),
			slog.String("method", resp.Request.Method),
			slog.String("url", resp.Request.URL.String()),
		}

		logLevel := slog.LevelInfo

		if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
			logLevel = slog.LevelDebug
			attrs = append(attrs, slog.String("body", string(respBody)))
		}

		slog.LogAttrs(context.Background(), logLevel, "complated request", attrs...)
	}

	return result
}

func ToJSONBodyBuffer[T any](input T) (*bytes.Buffer, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return &bytes.Buffer{}, err
	}

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

func ReqBodyToByte(req *http.Request) []byte {
	if req.Body == nil {
		return []byte{}
	}

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		return []byte{}
	}

	// should be re-assign into body
	req.Body = io.NopCloser(bytes.NewReader(reqBody))

	return reqBody
}

func ReqBodyToString(req *http.Request) string {
	reqBody := ReqBodyToByte(req)
	if reqBody == nil {
		return ""
	}

	return string(reqBody)
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

// Get RequestID from context
func GetRequestID(ctx context.Context) string {
	if val, ok := ctx.Value(requestIDKey).(string); ok {
		return val
	}
	return ""
}

// Get RequestID from context
func GetRequestIDKey() common.CtxKey {
	return requestIDKey
}
