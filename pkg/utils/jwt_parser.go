package utils

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-v2/pkg/configs"
	"strings"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Id                string
	UserID            int
	Expires           int64
	Username          string
	FirstName         string
	LastName          string
	Email             string
	MobilePhoneNumber string
}

type RefreshTokenMetaData struct {
	Id      string
	UserID  int
	Expires int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		//  ID.
		id := claims["id"].(string)
		// User ID.
		userID := int(claims["userId"].(float64))

		// Expires time.
		expires := int64(claims["expires"].(float64))

		meta := claims["meta_data"].(map[string]interface{})
		jsonData, _ := json.Marshal(meta)

		// Convert the JSON to a struct
		var structData *TokenMetadata
		json.Unmarshal(jsonData, &structData)

		res := structData
		res.Expires = expires
		res.Id = id
		res.UserID = userID

		return res, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(configs.Config.Apps.JwtSecretKey), nil
}

func ExtractRefreshTokenMetadata(tokenString string) (*RefreshTokenMetaData, error) {
	token, err := verifyRefreshToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		//  ID.
		id := claims["id"].(string)
		// User ID.
		userID := int(claims["userId"].(float64))

		// Expires time.
		expires := int64(claims["expires"].(float64))

		res := &RefreshTokenMetaData{
			Id:      id,
			UserID:  userID,
			Expires: expires,
		}
		return res, nil
	}

	return nil, err
}

func verifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, jwtRefreshKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtRefreshKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(configs.Config.Apps.JwtRefreshSecretKey), nil
}
