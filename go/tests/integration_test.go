package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type AddWeightRequest struct {
	Weight string `json:"weight"`
}

type WeightResponse struct {
	Weight    float64   `json:"weight"`
	Timestamp time.Time `json:"timestamp"`
}

type SuccessResponse[T any] struct {
	IsSuccess bool `json:"isSuccess"`
	Data      T    `json:"data,omitempty"`
}

type ErrorResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error,omitempty"`
}

func TestIntegration(t *testing.T) {
	baseURL, cleanup := SetupTestEnvironment(t)
	defer cleanup()

	var weights []WeightResponse

	t.Run("HealthCheck", func(t *testing.T) {
		var resp *http.Response
		var err error
		for range 10 {
			resp, err = http.Get(baseURL + "/health-check")
			if err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		require.NoError(t, err)
		require.NotNil(t, resp)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var healthResp SuccessResponse[string]
		err = json.Unmarshal(body, &healthResp)
		require.NoError(t, err)

		assert.True(t, healthResp.IsSuccess)
		assert.Equal(t, "OK", healthResp.Data)
	})

	t.Run("AddWeight", func(t *testing.T) {
		reqBody := AddWeightRequest{Weight: "82.5"}
		jsonBody, err := json.Marshal(reqBody)
		require.NoError(t, err)

		resp, err := http.Post(
			baseURL+"/weights",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		require.NoError(t, err)
		require.NotNil(t, resp)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var successResp SuccessResponse[string]
		err = json.Unmarshal(body, &successResp)
		require.NoError(t, err)

		assert.True(t, successResp.IsSuccess)
		assert.NotEmpty(t, successResp.Data)
	})

	t.Run("AddAnotherWeight", func(t *testing.T) {
		// Small delay to ensure different timestamps
		time.Sleep(100 * time.Millisecond)

		reqBody := AddWeightRequest{Weight: "83.1"}
		jsonBody, err := json.Marshal(reqBody)
		require.NoError(t, err)

		resp, err := http.Post(
			baseURL+"/weights",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		require.NoError(t, err)
		require.NotNil(t, resp)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var successResp SuccessResponse[string]
		err = json.Unmarshal(body, &successResp)
		require.NoError(t, err)

		assert.True(t, successResp.IsSuccess)
		assert.NotEmpty(t, successResp.Data)
	})

	t.Run("GetWeights", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/weights")
		require.NoError(t, err)
		require.NotNil(t, resp)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var successResp SuccessResponse[[]WeightResponse]
		err = json.Unmarshal(body, &successResp)
		require.NoError(t, err)

		assert.True(t, successResp.IsSuccess)

		weights = successResp.Data

		require.Len(t, weights, 2)

		assert.Equal(t, 82.5, weights[0].Weight)
		assert.False(t, weights[0].Timestamp.IsZero())

		assert.Equal(t, 83.1, weights[1].Weight)
		assert.False(t, weights[1].Timestamp.IsZero())

		assert.True(t, weights[0].Timestamp.Before(weights[1].Timestamp) ||
			weights[0].Timestamp.Equal(weights[1].Timestamp),
			"Weights should be sorted by timestamp ascending")
	})
}
