package nlp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	ExtractWord(_ *appcontext.AppContext, content string) (result *ExtractWordResult, err error)
	AnalyzeWord(_ *appcontext.AppContext, word, category string) (*AnalyzeWordResult, error)
}

type NLP struct {
	httpClient *resty.Client
}

func NewNLPClient(endpoint string) *NLP {
	return &NLP{
		httpClient: resty.New().
			SetBaseURL(endpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send NLP request at %s with status code %d", endpoint, resp.StatusCode())
			}),
	}
}
