package main

import (
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namhq1989/go-utilities/logger"
	"github.com/namhq1989/worddrop-server/internal/caching"
	"github.com/namhq1989/worddrop-server/internal/config"
	"github.com/namhq1989/worddrop-server/internal/database"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	appjwt "github.com/namhq1989/worddrop-server/internal/jwt"
	"github.com/namhq1989/worddrop-server/internal/monitoring"
	"github.com/namhq1989/worddrop-server/internal/monolith"
	"github.com/namhq1989/worddrop-server/internal/nlp"
	"github.com/namhq1989/worddrop-server/internal/queue"
	"github.com/namhq1989/worddrop-server/internal/tts"
	"github.com/namhq1989/worddrop-server/internal/utils/staticfiles"
	"github.com/namhq1989/worddrop-server/internal/utils/waiter"
	"time"
)

func main() {
	var err error

	// config
	cfg := config.Init()

	// logger
	logger.Init(cfg.Environment)

	// app error
	apperrors.Init()

	// static files
	staticfiles.Init(cfg.CDNEndpoint)

	// server
	a := app{}
	a.cfg = cfg

	// jwt
	a.jwt, err = appjwt.Init(cfg.AccessTokenSecret, time.Second*time.Duration(cfg.AccessTokenTTL))
	if err != nil {
		panic(err)
	}

	// rest
	a.rest = initRest(cfg)

	// grpc
	a.rpc = initRPC()

	// database
	a.database = database.NewDatabaseClient(cfg.PostgresConn)

	// caching
	a.caching = caching.NewCachingClient(cfg.CachingRedisURL)

	// tts
	a.tts = tts.NewTTSClient(tts.AwsConfig{
		AccessKey: cfg.AWSAccessKey,
		SecretKey: cfg.AWSSecretKey,
		Region:    cfg.AWSRegion,
	}, tts.R2Config{
		AccessKey: cfg.R2AccessKey,
		SecretKey: cfg.R2SecretKey,
		Endpoint:  cfg.R2Endpoint,
		Bucket:    cfg.R2Bucket,
	})

	// nlp
	a.nlp = nlp.NewNLPClient(cfg.NLPEndpoint)

	// monitoring
	a.monitoring = monitoring.NewMonitoringClient(
		a.rest,
		monitoring.OtelConfig{
			Endpoint:   cfg.OpenObserveHttpEndpoint,
			StreamName: cfg.OpenObserveStreamName,
			Token:      cfg.OpenObserveToken,
		},
		monitoring.SentryConfig{
			Dsn:         cfg.SentryDSN,
			MachineName: cfg.SentryMachineName,
		},
		cfg.AppName,
		cfg.Environment,
	)

	// queue
	a.queue = queue.Init(cfg.QueueRedisURL, cfg.QueueConcurrency)

	// init queue's dashboard
	a.rest.Any(fmt.Sprintf("%s/*", queue.DashboardPath), echo.WrapHandler(queue.EnableDashboard(cfg.QueueRedisURL)), middleware.BasicAuth(func(username, password string, _ echo.Context) (bool, error) {
		if !cfg.IsEnvRelease {
			return true, nil
		}
		return subtle.ConstantTimeCompare([]byte(username), []byte(cfg.QueueUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.QueuePassword)) == 1, nil
	}))

	// waiter
	a.waiter = waiter.New(waiter.CatchSignals())

	// modules
	a.modules = []monolith.Module{}

	// start
	if err = a.startupModules(); err != nil {
		panic(err)
	}

	fmt.Println("--- started worddrop-server application")
	defer fmt.Println("--- stopped worddrop-server application")

	// wait for other service starts
	a.waiter.Add(
		a.waitForRest,
		a.waitForRPC,
	)
	if err = a.waiter.Wait(); err != nil {
		panic(err)
	}
}
