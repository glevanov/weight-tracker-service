package tests

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	TestDBName = "weight-tracker-test"
)

func SetupTestEnvironment(t *testing.T) (string, string, func()) {
	ctx := context.Background()

	// Create a Docker network for container communication
	nw, err := network.New(ctx)
	require.NoError(t, err, "Failed to create test network")
	require.NotNil(t, nw, "Network should not be nil")

	mongoContainer, err := setupMongoContainer(t, nw.Name)
	require.NoError(t, err, "Failed to start MongoDB container")

	// Get MongoDB connection URI for test setup
	mongoHost, err := mongoContainer.Host(ctx)
	require.NoError(t, err, "Failed to get MongoDB host")

	mongoPort, err := mongoContainer.MappedPort(ctx, "27017")
	require.NoError(t, err, "Failed to get MongoDB port")

	mongoURI := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort.Port())

	baseURL, serviceCleanup := setupServiceContainer(t, nw.Name)

	cleanup := func() {
		serviceCleanup()
		if err := mongoContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate MongoDB container: %v", err)
		}
		if err := nw.Remove(ctx); err != nil {
			t.Logf("Failed to remove test network: %v", err)
		}
	}

	return baseURL, mongoURI, cleanup
}

func setupMongoContainer(t *testing.T, networkName string) (testcontainers.Container, error) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:7",
		ExposedPorts: []string{"27017/tcp"},
		Networks:     []string{networkName},
		NetworkAliases: map[string][]string{
			networkName: {"mongo"},
		},
		WaitingFor: wait.ForListeningPort("27017/tcp").WithStartupTimeout(60 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

func setupServiceContainer(t *testing.T, networkName string) (string, func()) {
	ctx := context.Background()

	projectRoot, err := filepath.Abs("..")
	require.NoError(t, err, "Failed to get project root")

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    projectRoot,
			Dockerfile: "Dockerfile.test",
		},
		ExposedPorts: []string{"3000/tcp"},
		Networks:     []string{networkName},
		WaitingFor:   wait.ForListeningPort("3000/tcp").WithStartupTimeout(60 * time.Second),
		Env: map[string]string{
			"PORT":           "3000",
			"FRONTEND_URL":   "http://localhost:5173",
			"CONNECTION_URI": "mongodb://mongo:27017",
			"DB_NAME":        TestDBName,
			"JWT_SECRET":     "test-jwt-secret-key",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	require.NoError(t, err, "Failed to start service container")
	require.NotNil(t, container, "Container should not be nil")

	mappedPort, err := container.MappedPort(ctx, "3000")
	require.NoError(t, err, "Failed to get mapped port")

	host, err := container.Host(ctx)
	require.NoError(t, err, "Failed to get container host")

	baseURL := fmt.Sprintf("http://%s:%s", host, mappedPort.Port())

	cleanup := func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate service container: %v", err)
		}
	}

	return baseURL, cleanup
}
