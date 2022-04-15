package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	CompanyServiceHost string
	CompanyServicePort int

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string
	MinioBucketName  string
	MinioLocation    string
	MinioHost        string

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	config.CompanyServiceHost = cast.ToString(getOrReturnDefault("COMPANY_SERVICE_HOST", "localhost"))
	config.CompanyServicePort = cast.ToInt(getOrReturnDefault("COMPANY_SERVICE_PORT", 5004))

	config.MinioEndpoint = cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "test.cdn.urecruit.udevs.io"))
	config.MinioAccessKeyID = cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "2R5YabYDYwesXPDPprWc6DpbczCsXL97"))
	config.MinioSecretKey = cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY_ID", "Ps5Che6XtJ6JmvsFXrXUH3tnhxwnZNYh"))
	config.MinioBucketName = cast.ToString(getOrReturnDefault("MINIO_BACKET_NAME", "photos"))
	config.MinioLocation = cast.ToString(getOrReturnDefault("MINIO_LOCATION", "us-east-1"))
	config.MinioHost = cast.ToString(getOrReturnDefault("MINIO_HOST", "test.cdn.urecruit.udevs.io"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
