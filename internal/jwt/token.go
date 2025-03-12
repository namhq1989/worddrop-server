package appjwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/go-utilities/uuid"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
)

var unauthorizedError = errors.New("unauthorized")

func (j JWT) GenerateAccessToken(ctx *appcontext.AppContext, userID, platformID string, province int) (string, error) {
	if !uuid.IsValidID(userID) || !uuid.IsValidID(platformID) {
		ctx.Logger().Error("invalid user id or platform id", nil, appcontext.Fields{"userID": userID, "platformID": platformID})
		return "", apperrors.Common.InvalidIdentifier
	}

	accessToken, _, err := j.generateAccessToken(userID, platformID, province)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{"userID": userID, "platformID": platformID})
		return "", err
	}

	return accessToken, nil
}

func (j JWT) generateAccessToken(userID, platformID string, province int) (string, time.Time, error) {
	exp := time.Now().Add(j.accessTokenTTL)
	claims := &Claims{
		UserID:     userID,
		PlatformID: platformID,
		Province:   province,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	value, err := token.SignedString(j.accessTokenSecret)
	return value, exp, err
}

func (j JWT) ParseAccessToken(ctx *appcontext.AppContext, token string) (*Claims, error) {
	if token == "" {
		return nil, unauthorizedError
	}

	// parse the token
	tokenData, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			ctx.Logger().Error("check signing method", fmt.Errorf("unexpected signing method: %v", t.Header["alg"]), appcontext.Fields{"token": token})
			return nil, unauthorizedError
		}

		return j.accessTokenSecret, nil
	})

	// error
	if err != nil {
		ctx.Logger().Error("parse token", err, appcontext.Fields{"token": token})
		return nil, err
	}

	// respond
	if claims, ok := tokenData.Claims.(*Claims); ok && tokenData.Valid {
		return claims, nil
	} else {
		ctx.Logger().Error("parse claims", nil, appcontext.Fields{"token": token, "tokenData": tokenData})
		return nil, unauthorizedError
	}
}
