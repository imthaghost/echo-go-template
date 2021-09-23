package middleware

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

// UserPayload is context of decoded verified token claims
type UserPayload struct {
	ID    string // Unique ID
}

// TokenVerification ...
func TokenVerification(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get auth token from header
		bearer := c.Request().Header.Get("Authorization")
		items := strings.Fields(bearer)

		// check
		if len(items) < 2 {
			return echo.NewHTTPError(http.StatusForbidden, "Bearer: Authorization token not found")
		}

		// token
		token := items[1]
		log.Print(token)
		// parse and verify
		//t, err := ParseAndVerifyJWT(token)
		//if err != nil {
		//	return echo.NewHTTPError(http.StatusForbidden, err)
		//}

		// get decoded claims
		//claims, ok := t.Claims.(jw)
		//if !ok {
		//	return echo.NewHTTPError(http.StatusBadRequest, "Could not extract token claims")
		//}

		// create payload
		//fullName := claims["name"].(string)
		//phoneNumber := claims["phone_number"].(string)
		//ID := claims["sub"].(string)
		payload := UserPayload{ID: ""}


		// set payload context
		c.Set("Payload", payload)

		return next(c)
	}
}