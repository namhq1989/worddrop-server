package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/utils/httprespond"
	"github.com/namhq1989/worddrop-server/internal/utils/validation"
	"github.com/namhq1989/worddrop-server/pkg/content/dto"
)

func (s server) registerWordRoutes() {
	g := s.echo.Group("/api/word")

	g.GET("/initial", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.GetInitialWordsRequest)
		)

		resp, err := s.app.GetInitialWords(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetInitialWordsRequest](next)
	})

	g.GET("/new", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.GetNewWordRequest)
		)

		resp, err := s.app.GetNewWord(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetNewWordRequest](next)
	})
}
