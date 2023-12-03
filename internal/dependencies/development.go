//go:build dev || e2e
// +build dev e2e

package dependencies

import (
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/testcontainers/testcontainers-go"
)

// init will be used to start up the containers for development mode. It will use
// testcontainers-go to start up the following containers:
// - TBA
// Please read this blog post for more information: https://www.atomicjar.com/2023/08/local-development-of-go-applications-with-testcontainers/
func init() {
	startupDependenciesFns := []func() (testcontainers.Container, error){
		startDatabase,
	}

	runtimeDependencies := make([]testcontainers.Container, 0, len(startupDependenciesFns))

	for _, fn := range startupDependenciesFns {
		c, err := fn()
		if err != nil {
			panic(err)
		}
		runtimeDependencies = append(runtimeDependencies, c)
	}

	// register a graceful shutdown to stop the dependencies when the application is stopped
	// only in development mode
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGKILL)
	go func() {
		// also use the shutdown function when the SIGTERM or SIGINT signals are received
		sig := <-gracefulStop
		log.Printf("caught sig: %+v\n", sig)
		err := shutdownDependencies(runtimeDependencies...)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()
}

// helper function to stop the dependencies
func shutdownDependencies(containers ...testcontainers.Container) error {
	for _, c := range containers {
		err := c.Terminate(ctx)
		if err != nil {
			return fmt.Errorf("failed to terminate container: %w", err)
		}
	}

	return nil
}

func startDatabase() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Name: "astroboy-db-local",
		// we're using the latest version of the image provided by Amazon
		Image: "amazon/dynamodb-local:latest",
		// be sure to use the commands as described in the documentation, but
		// an in-memory version is good enough for us
		Cmd: []string{"-jar", "DynamoDBLocal.jar", "-inMemory"},
		// by default, DynamoDB runs on port 8000
		ExposedPorts: []string{"8000/tcp"},
		// testcontainers let's us block until the port is available, i.e.,
		// DynamoDB has started
		WaitingFor: wait.NewHostPortStrategy("8000"),
	}

	// let's start the container!
	d, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            true,
	})

	if err != nil {
		return nil, err
	}

	host, err := d.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := d.MappedPort(ctx, "8000")
	if err != nil {
		return nil, err
	}

	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port.Port())

	return d, nil
}
