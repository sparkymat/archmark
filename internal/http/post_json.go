package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	httplib "net/http"
	"time"
)

const HTTPTimeoutSeconds = 20

var ErrRequestFailed = errors.New("request failed")

func PostJSON[R any, Q any](ctx context.Context, httpClient Client, url string, headers map[string]string, req Q) (R, error) {
	var response R

	ctx, cancel := context.WithTimeout(ctx, HTTPTimeoutSeconds*time.Second)
	defer cancel()

	requestBytes, err := json.Marshal(req)
	if err != nil {
		return response, fmt.Errorf("unable to marshal request. err: %w", err)
	}

	requestBody := bytes.NewReader(requestBytes)

	httpRequest, err := httplib.NewRequestWithContext(ctx, httplib.MethodPost, url, requestBody)
	if err != nil {
		return response, fmt.Errorf("unable to form request. err: %w", err)
	}

	for k, v := range headers {
		httpRequest.Header.Set(k, v)
	}

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return response, fmt.Errorf("request failed. err: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("failed to read response. err: %w", err)
	}

	if resp.StatusCode != httplib.StatusOK {
		return response, fmt.Errorf("request failed with %d. err: %w", resp.StatusCode, ErrRequestFailed)
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, fmt.Errorf("failed to unmarshal response. err: %w", err)
	}

	return response, nil
}
