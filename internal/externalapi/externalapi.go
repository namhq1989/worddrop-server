package externalapi

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	FindWordWitDatamuse(ctx *appcontext.AppContext, word string) (*DatamuseFindWordResult, error)
	FetchNews(ctx *appcontext.AppContext) (*FetchNewsResult, error)
}

type ExternalAPI struct {
	datamuse     *resty.Client
	rapidApiNews *resty.Client
}

const (
	datamuseApiEndpoint  = "https://api.datamuse.com"
	rapidApiNewsEndpoint = "https://news67.p.rapidapi.com"
)

func NewExternalAPIClient(rapidAPIKey string) *ExternalAPI {
	return &ExternalAPI{
		datamuse: resty.New().
			SetBaseURL(datamuseApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Datamuse request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
		rapidApiNews: resty.New().
			SetBaseURL(rapidApiNewsEndpoint).
			SetHeader("Accept", "application/json").
			SetHeader("x-rapidapi-host", "news67.p.rapidapi.com").
			SetHeader("x-rapidapi-key", "rapidAPIKey").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send RapidAPI request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
