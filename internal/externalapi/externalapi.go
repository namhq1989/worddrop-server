package externalapi

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	FetchNewsWithNewsService(ctx *appcontext.AppContext) (*FetchNewsWithNewsServiceResult, error)
	FetchNewsWithGoogleService(ctx *appcontext.AppContext) (*FetchNewsWithGoogleServiceResult, error)
}

type ExternalAPI struct {
	rapidApiKeys       []string
	rapidApiKeysLength int
	rapidApiNews       *resty.Client
	rapidApiGoogleNews *resty.Client
}

const (
	rapidApiNewsEndpoint       = "https://news67.p.rapidapi.com"
	rapidApiGoogleNewsEndpoint = "https://google-api31.p.rapidapi.com"
)

func NewExternalAPIClient(rapidAPIKeys []string) *ExternalAPI {
	return &ExternalAPI{
		rapidApiKeys:       rapidAPIKeys,
		rapidApiKeysLength: len(rapidAPIKeys),
		rapidApiNews: resty.New().
			SetBaseURL(rapidApiNewsEndpoint).
			SetHeader("Accept", "application/json").
			SetHeader("x-rapidapi-host", "news67.p.rapidapi.com").
			SetTimeout(60 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send RapidAPI News request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
		rapidApiGoogleNews: resty.New().
			SetBaseURL(rapidApiGoogleNewsEndpoint).
			SetHeader("Accept", "application/json").
			SetHeader("x-rapidapi-host", "google-api31.p.rapidapi.com").
			SetTimeout(60 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send RapidAPI Google News request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
