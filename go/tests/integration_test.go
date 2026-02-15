package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type HealthResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Data      string `json:"data"`
}

func TestIntegration(t *testing.T) {
	baseURL, cleanup := setupServiceContainer(t)
	defer cleanup()

	var resp *http.Response
	var err error
	for i := 0; i < 5; i++ {
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

	var healthResp HealthResponse
	err = json.Unmarshal(body, &healthResp)
	require.NoError(t, err)

	assert.True(t, healthResp.IsSuccess)
	assert.Equal(t, "OK", healthResp.Data)
}
