package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-v2/pkg/configs"
	"strconv"
	"strings"
	"time"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access   string
	Refresh  string
	ExpireIn int64
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id string, metaData TokenMetadata) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, e, err := generateNewAccessToken(id, metaData)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateRefreshToken(id, metaData.UserID)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:   accessToken,
		Refresh:  refreshToken,
		ExpireIn: e,
	}, nil
}

func generateNewAccessToken(id string, metaData TokenMetadata) (string, int64, error) {
	// Set secret key from .env file.
	secret := configs.Config.Apps.JwtSecretKey

	// Set expires minutes count for secret key from resource file.
	minutesCount := configs.Config.Apps.TokenExpiration

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["userId"] = metaData.UserID
	claims["expires"] = time.Now().Add(time.Second * time.Duration(minutesCount)).Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(minutesCount)).Unix()
	claims["meta_data"] = metaData

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", 0, err
	}

	return t, minutesCount, nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := configs.Config.Apps.JwtRefreshSecretKey + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expires hours count for refresh key from resource file.
	hoursCount := configs.Config.Apps.RefreshTokenExpiration

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

func generateRefreshToken(id string, userId int) (string, error) {
	// Set secret key from .env file.
	secret := configs.Config.Apps.JwtRefreshSecretKey

	// Set expires minutes count for secret key from resource file.
	minutesCount := configs.Config.Apps.RefreshTokenExpiration

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["userId"] = userId
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
