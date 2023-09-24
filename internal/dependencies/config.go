package dependencies

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	AwsSqsEndpoint  string `env:"AWS_SQS_ENDPOINT"`
	AwsRegion       string `env:"AWS_REGION"`
	SqsQueueUrl     string `env:"SQS_QUEUE_URL"`
	KafkaBrokerAddr string `env:"KAFKA_BROKER_ADDR"`
	KafkaTopic      string `env:"KAFKA_TOPIC"`
	RedisAddr       string `env:"REDIS_ADDR"`
}

func LoadEnv() *Config {
	appEnv := os.Getenv("APP_ENV")

	if "" == appEnv {
		appEnv = "dev"
	}

	log.Printf("Loading %s config\n", appEnv)
	err := godotenv.Load(dir(".env." + appEnv))
	if err != nil {
		log.Fatalf("error loading app config: %v", err.Error())
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error loading app config: %v", err.Error())
	}

	return &cfg
}

// dir returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
