package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID      uuid.UUID
	Credentials map[string]bool
	Expires     int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *echo.Context) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		// Expires time.
		expires := int64(claims["expires"].(float64))

		// User credentials.
		credentials := map[string]bool{
			"book:create": claims["book:create"].(bool),
			"book:update": claims["book:update"].(bool),
			"book:delete": claims["book:delete"].(bool),
		}

		return &TokenMetadata{
			UserID:      userID,
			Credentials: credentials,
			Expires:     expires,
		}, nil
	}

	return nil, err
}

