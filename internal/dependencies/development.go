package dependencies

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
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
	go func() {
		// also use the shutdown function when the SIGTERM or SIGINT signals are received
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\n", sig)
		//err := shutdownDependencies(runtimeDependencies...)
		//if err != nil {
		//	os.Exit(1)
		//}
		os.Exit(0)
	}()
}

func startDatabase() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
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
	})

	if err != nil {
		panic(err)
	}

	defer d.Terminate(context.Background())

	return nil, nil
}
