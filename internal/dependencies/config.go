package dependencies

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	AwsRegion     string `env:"AWS_REGION" envDefault:"eu-west-1"`
	RedisAddr     string `env:"REDIS_ADDR"`
	WebPort       string `env:"PORT" envDefault:"5000"`
	DevMode       bool   `env:"DEV_MODE" envDefault:"false"`
	DbTableName   string `env:"DB_TABLE_NAME" envDefault:"astroboy-store"`
	DbHost        string `env:"DB_HOST" envDefault:"localhost"`
	DbPort        string `env:"DB_PORT" envDefault:"8000"`
	JwtSigningKey string `env:"JWT_SIGNING_KEY" envDefault:"secret"`
}

func LoadEnv() *Config {
	appEnv := os.Getenv("APP_ENV")
	cfg := Config{}

	if "" != appEnv {
		log.Printf("Loading %s config\n", appEnv)
		err := godotenv.Load(dir(".env." + appEnv))
		if err != nil {
			log.Fatalf("error loading app config: %v", err.Error())
		}

		cfg.DevMode = true
	}

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
