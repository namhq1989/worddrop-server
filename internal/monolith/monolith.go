package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/caching"
	"github.com/namhq1989/worddrop-server/internal/config"
	"github.com/namhq1989/worddrop-server/internal/database"
	"github.com/namhq1989/worddrop-server/internal/externalapi"
	appjwt "github.com/namhq1989/worddrop-server/internal/jwt"
	"github.com/namhq1989/worddrop-server/internal/monitoring"
	"github.com/namhq1989/worddrop-server/internal/nlp"
	"github.com/namhq1989/worddrop-server/internal/queue"
	"github.com/namhq1989/worddrop-server/internal/tts"
	"github.com/namhq1989/worddrop-server/internal/utils/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Caching() *caching.Caching
	JWT() *appjwt.JWT
	Queue() *queue.Queue
	TTS() *tts.TTS
	NLP() *nlp.NLP
	ExternalAPI() *externalapi.ExternalAPI
	Monitoring() *monitoring.Monitoring
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
