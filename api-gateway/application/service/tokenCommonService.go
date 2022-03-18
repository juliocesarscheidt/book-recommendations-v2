package service

import (
	"errors"
	"strings"
	"io/ioutil"
	"crypto/rsa"
	"path/filepath"

	jwt "github.com/golang-jwt/jwt/v4"
)

func LoadPublicKeyRSA(location string) (*rsa.PublicKey, error) {
	absLocation, _ := filepath.Abs(location)
	keyData, err := ioutil.ReadFile(absLocation)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func ExtractTokenString(authHeader string) (string, error) {
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	if tokenString == "" {
		return "", errors.New("Invalid Authorization Header")
	}
	return tokenString, nil
}

func ParseTokenUsingPublicKey(tokenString string) (*jwt.Token, error) {
	publicKey, err := LoadPublicKeyRSA("keys/jwtencyptionkey.pub")
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Invalid Signing Method")
		}
		return publicKey, nil
	})
	return token, err
}

func ValidateClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if (! ok || ! token.Valid) {
		return nil, errors.New("Invalid Claims")
	}

	expectedAudience := "urn:book-recommendations:api-gateway"
	expectedIssuer := "urn:book-recommendations:users-microservice"
	if (claims["aud"] != expectedAudience || claims["iss"] != expectedIssuer) {
		return nil, errors.New("Invalid Audience or Issuer")
	}

	return claims, nil
}
