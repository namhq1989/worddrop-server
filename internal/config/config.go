package config

import "errors"

type (
	Server struct {
		RestPort string
		GRPCPort string

		AppName      string
		Environment  string
		IsEnvRelease bool

		// Authentication
		FirebaseServiceAccount      string
		AnonymousUserChecksumSecret string
		AccessTokenSecret           string
		AccessTokenTTL              int // seconds

		// Postgres
		PostgresConn string

		// Redis
		CachingRedisURL string

		// Queue
		QueueRedisURL    string
		QueueUsername    string
		QueuePassword    string
		QueueConcurrency int

		// Open Observe
		OpenObserveHttpEndpoint string
		OpenObserveStreamName   string
		OpenObserveToken        string

		// Sentry
		SentryDSN         string
		SentryMachineName string

		// Endpoint
		NLPEndpoint string
		CDNEndpoint string
	}
)

func Init() Server {
	cfg := Server{
		RestPort: ":3000",
		GRPCPort: ":3001",

		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),

		FirebaseServiceAccount:      getEnvStr("FIREBASE_SERVICE_ACCOUNT"),
		AnonymousUserChecksumSecret: getEnvStr("ANONYMOUS_USER_CHECKSUM_SECRET"),
		AccessTokenSecret:           getEnvStr("ACCESS_TOKEN_SECRET"),
		AccessTokenTTL:              getEnvInt("ACCESS_TOKEN_TTL"),

		PostgresConn: getEnvStr("POSTGRES_CONN"),

		CachingRedisURL: getEnvStr("CACHING_REDIS_URL"),

		QueueRedisURL:    getEnvStr("QUEUE_REDIS_URL"),
		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),

		OpenObserveHttpEndpoint: getEnvStr("OPEN_OBSERVE_HTTP_ENDPOINT"),
		OpenObserveStreamName:   getEnvStr("OPEN_OBSERVE_STREAM_NAME"),
		OpenObserveToken:        getEnvStr("OPEN_OBSERVE_TOKEN"),

		SentryDSN:         getEnvStr("SENTRY_DSN"),
		SentryMachineName: getEnvStr("SENTRY_MACHINE_NAME"),

		NLPEndpoint: getEnvStr("NLP_ENDPOINT"),
		CDNEndpoint: getEnvStr("CDN_ENDPOINT"),
	}
	cfg.IsEnvRelease = cfg.Environment == "release"

	// validation
	if cfg.Environment == "" {
		panic(errors.New("missing ENVIRONMENT"))
	}

	if cfg.FirebaseServiceAccount == "" {
		panic(errors.New("missing FIREBASE_SERVICE_ACCOUNT"))
	}

	if cfg.AccessTokenSecret == "" {
		panic(errors.New("missing ACCESS_TOKEN_SECRET"))
	}

	if cfg.PostgresConn == "" {
		panic(errors.New("missing POSTGRES_CONN"))
	}

	if cfg.CachingRedisURL == "" {
		panic(errors.New("missing CACHING_REDIS_URL"))
	}

	if cfg.QueueRedisURL == "" {
		panic(errors.New("missing QUEUE_REDIS_URL"))
	}

	if cfg.CDNEndpoint == "" {
		panic(errors.New("missing CDN_ENDPOINT"))
	}

	return cfg
}
