package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-http-utils/headers"
	"github.com/golang-jwt/jwt"
	"github.com/willkoerich/rock-challenge/cmd/api/response"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"net/http"
	"strings"
	"time"
)

const (
	ErrorToDecodeSubmittedLoginInformationMessage = "unable to decode submitted login information body."
	InvalidLoginInformationMessage                = "unable to authenticate with submitted credentials."
	MissingAuthorizationHeaderMessage             = "unauthorized due to missing authorization token"
	UnauthorizedMessage                           = "unauthorized"
	ErrorToParseJWTAccessTokenMessage             = "unauthorized due to error parsing the JWT"
	InvalidJWTAccessTokenMessage                  = "unauthorized due to invalid token"
)

type (
	LoginHandler struct {
		controller domain.LoginController
	}
)

var (
	appKey = "stone"
)

func NewLoginHandler(controller domain.LoginController) LoginHandler {
	return LoginHandler{
		controller: controller,
	}
}

func (handler LoginHandler) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var decodedBody domain.AuthenticationRequest
	if err := json.NewDecoder(request.Body).Decode(&decodedBody); err != nil {
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			ErrorToDecodeSubmittedLoginInformationMessage, &err)
		return
	}
	authenticationResponse, err := handler.controller.Authenticate(request.Context(), decodedBody)
	if err != nil {
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			InvalidLoginInformationMessage, &err)
		return
	}

	tokenString, err := generateToken(authenticationResponse.AccountID, appKey)
	if err != nil {
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			InvalidLoginInformationMessage, &err)
		return
	}

	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(domain.TokenGenerationResponse{AccessToken: tokenString})
	if err != nil {
		return
	}
	return
}

func (handler LoginHandler) VerifyToken(writer http.ResponseWriter, request *http.Request) (domain.AuthenticationResponse, error) {
	if request.Header[headers.Authorization] != nil {
		return validateToken(writer, request, appKey)
	}
	return domain.AuthenticationResponse{}, errors.New(MissingAuthorizationHeaderMessage)
}

func generateToken(accountID int, appKey string) (string, error) {
	expirationTime := float64(time.Now().Add(10 * time.Minute).Unix())

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expirationTime
	claims["authorized"] = true
	claims["user"] = accountID

	return token.SignedString([]byte(appKey))
}

func validateToken(writer http.ResponseWriter, request *http.Request, key string) (domain.AuthenticationResponse, error) {
	account := domain.AuthenticationResponse{}
	extractedToken, err := extractAuthorizationTokenValue(request)
	if err != nil {
		return domain.AuthenticationResponse{}, err
	}

	token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			response.CreateHandlerResponse(writer, http.StatusBadRequest,
				UnauthorizedMessage, &err)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			account = domain.AuthenticationResponse{AccountID: int(claims["user"].(float64))}
		}

		return []byte(key), nil
	})

	if err != nil {
		return domain.AuthenticationResponse{}, errors.New(ErrorToParseJWTAccessTokenMessage)
	}

	if !token.Valid {
		return domain.AuthenticationResponse{}, errors.New(InvalidJWTAccessTokenMessage)
	}
	return account, nil
}

func extractAuthorizationTokenValue(request *http.Request) (string, error) {
	if len(request.Header[headers.Authorization][0]) > 6 &&
		strings.ToUpper(request.Header[headers.Authorization][0][0:7]) == "BEARER " {
		return request.Header[headers.Authorization][0][7:], nil
	}
	return "", errors.New(MissingAuthorizationHeaderMessage)
}
