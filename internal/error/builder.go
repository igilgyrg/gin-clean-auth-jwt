package error

import (
	"errors"
	"net/http"
)

var (
	Unauthorized          = errors.New("unauthorized")
	BadRequest            = errors.New("bad request")
	InternalSystem        = errors.New("internal system error")
	ItemIsExists          = errors.New("item is exists")
	InvalidRefreshToken   = errors.New("invalid refresh token")
	InvalidAccessToken    = errors.New("invalid access token")
	UserIsExistsWithEmail = errors.New("user is exists with email")
)

func NewUnauthorizedError(causes interface{}) Error {
	return &HttpError{
		HttpStatus: http.StatusUnauthorized,
		HttpError:  Unauthorized.Error(),
		HttpCauses: causes,
	}
}

func NewBadRequestError(causes interface{}, msg string) Error {
	return &HttpError{
		HttpStatus:  http.StatusBadRequest,
		HttpError:   BadRequest.Error(),
		HttpMessage: msg,
		HttpCauses:  causes,
	}
}

func NewUserIsExistsWithEmail(causes interface{}, msg string) Error {
	return &HttpError{
		HttpStatus:  http.StatusBadRequest,
		HttpError:   UserIsExistsWithEmail.Error(),
		HttpMessage: msg,
		HttpCauses:  causes,
	}
}

func NewRefreshTokenInvalid(causes interface{}, msg string) Error {
	return &HttpError{
		HttpStatus:  http.StatusUnauthorized,
		HttpError:   InvalidRefreshToken.Error(),
		HttpMessage: msg,
		HttpCauses:  causes,
	}
}

func NewAccessTokenInvalid(causes interface{}, msg string) Error {
	return &HttpError{
		HttpStatus:  http.StatusUnauthorized,
		HttpError:   InvalidAccessToken.Error(),
		HttpMessage: msg,
		HttpCauses:  causes,
	}
}
