package tests

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupServiceContainer(t *testing.T) (string, func()) {
	ctx := context.Background()

	projectRoot, err := filepath.Abs("..")
	require.NoError(t, err, "Failed to get project root")

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    projectRoot,
			Dockerfile: "Dockerfile.test",
		},
		ExposedPorts: []string{"3000/tcp"},
		WaitingFor:   wait.ForListeningPort("3000/tcp").WithStartupTimeout(60 * time.Second),
		Env: map[string]string{
			"PORT":         "3000",
			"FRONTEND_URL": "http://localhost:5173",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	require.NoError(t, err, "Failed to start container")
	require.NotNil(t, container, "Container should not be nil")

	mappedPort, err := container.MappedPort(ctx, "3000")
	require.NoError(t, err, "Failed to get mapped port")

	host, err := container.Host(ctx)
	require.NoError(t, err, "Failed to get container host")

	baseURL := fmt.Sprintf("http://%s:%s", host, mappedPort.Port())

	cleanup := func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}

	return baseURL, cleanup
}
