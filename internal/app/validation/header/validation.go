package header

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/common-go/jwt"
)

var (
	SecretKey          = "SECRETKEY"
	errNoTokenInHeader = errors.New("no authentication token in header")
	errInvalidToken    = errors.New("invalid token")
)

const (
	// RequestIDLimit is max
	RequestIDLimit = 50

	// RequestIDLimitMin is min
	RequestIDLimitMin = 1

	// AppIDLimit is appID limit
	AppIDLimit = 5

	// DateTimeLimit is datetime limit
	DateTimeLimit = 43

	// OriginalUIDLimit is original uid limit
	OriginalUIDLimit = 50

	TokenHeaderKey = "token"

	AuthHeaderKey = "authorization"

	BearerPrefix = "Bearer "
)

// isValidRequestID func
func isValidRequestID(rqID string) error {
	if len(rqID) > RequestIDLimit || len(rqID) == RequestIDLimitMin {
		return errors.New("request_id is out of length")
	}
	if IsEmpty(rqID) {
		return errors.New("request_id is missing")
	}
	return nil
}

// isValidAppID func
func isValidAppID(rqID string) error {
	if len(rqID) > AppIDLimit {
		return errors.New("request_app_id is out of length")
	}
	if IsEmpty(rqID) {
		return errors.New("request_app_id is missing")
	}
	return nil
}

// isValidRequestDate func
func isValidRequestDate(date *time.Time) error {
	if date == nil || date.IsZero() {
		return errors.New("request_datetime is missing")
	}
	if len(date.String()) > DateTimeLimit {
		return errors.New("request_datetime is out of length")
	}
	now := time.Now()
	if !(date.Year() == now.Year() && date.Month() == now.Month() && date.Day() == now.Day()) {
		return errors.New("invalid request_date")
	}
	return nil
}

// IsEmpty func
func IsEmpty(field string) bool {
	return len(field) == 0
}

//GetTokenFromHttpRequest func
func GetTokenFromHttpRequest(r *http.Request) (string, error) {
	if authString := r.Header.Get(AuthHeaderKey); !strings.HasPrefix(authString, BearerPrefix) {
		return "", errNoTokenInHeader

	} else {
		return authString[len(BearerPrefix):], nil
	}
}

// ValidateJWT func
func ValidateJWT(context context.Context) error {
	tokenService := jwt.DefaultTokenService{}
	tokenString := context.Value(TokenHeaderKey).(string)
	_, _, _, err := tokenService.VerifyToken(tokenString, SecretKey)
	if err != nil {
		return err
	}
	return nil
}
