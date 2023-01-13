package handlers

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"net/http"
	"strings"
	"time"
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
		writer.WriteHeader(http.StatusBadRequest)
		_, err2 := writer.Write([]byte("unable to decode submitted login information body. Error: " + err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	authenticationResponse, err := handler.controller.Authenticate(request.Context(), decodedBody)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_, err := writer.Write([]byte("unable to authenticate with submitted credentials."))
		if err != nil {
			return
		}
		return
	}

	tokenString, err := generateToken(authenticationResponse.AccountID, appKey)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte(tokenString))
	if err != nil {
		return
	}
}

func (handler LoginHandler) VerifyToken(writer http.ResponseWriter, request *http.Request) (domain.Account, error) {
	if request.Header["Authorization"] != nil {
		return validateToken(writer, request, appKey)
	}
	return domain.Account{}, errors.New("unauthorized due to missing authorization token")
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

func validateToken(writer http.ResponseWriter, request *http.Request, key string) (domain.Account, error) {
	account := domain.Account{}
	extractedToken, err := extractAuthorizationTokenValue(request)
	if err != nil {
		return domain.Account{}, err
	}

	token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("unauthorized"))
			if err != nil {
				return domain.Account{}, err
			}
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			account = domain.Account{ID: claims["user"].(int)}
		}

		return []byte(key), nil
	})

	if err != nil {
		return domain.Account{}, errors.New("unauthorized due to error parsing the JWT")
	}

	if !token.Valid {
		return domain.Account{}, errors.New("unauthorized due to invalid token")
	}
	return account, nil
}

func extractAuthorizationTokenValue(request *http.Request) (string, error) {
	if len(request.Header["Authorization"][0]) > 6 &&
		strings.ToUpper(request.Header["Authorization"][0][0:7]) == "BEARER " {
		return request.Header["Authorization"][0][7:], nil
	}
	return "", errors.New("unauthorized due to missing authorization header value")
}
