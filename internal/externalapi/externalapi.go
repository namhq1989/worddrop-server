package externalapi

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	FetchNews(ctx *appcontext.AppContext) (*FetchNewsResult, error)
}

type ExternalAPI struct {
	rapidApiKeys       []string
	rapidApiKeysLength int
	rapidApiNews       *resty.Client
}

const (
	rapidApiNewsEndpoint = "https://news67.p.rapidapi.com"
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
				return 1, fmt.Errorf("failed to send RapidAPI request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
