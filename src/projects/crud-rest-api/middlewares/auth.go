package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	models "../models"
	u "../utils"
	"github.com/dgrijalva/jwt-go"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// List of endpoints that do not require Auth
		unAuthEndpoints := []string{
			"/api/v1/auth/login",
			"/api/v1/auth/register",
			"/api/v1/posts",
		}
		// Current request path
		urlPath := req.URL.Path
		// Check if the endpoint is among the list of unauthenticated endpoints
		// If yes, serve the endpoint
		for _, value := range unAuthEndpoints {
			if value == urlPath {
				next.ServeHTTP(w, req)
				return
			}
		}

		token := req.Header.Get("Authorization")
		var response map[string]interface{}

		// If auth token is not included in the Header, return error message with Forbidden status code
		if token == "" {
			response = u.Message(false, "Authorization token header is required.")
			u.BuildResponse(w, 403, response)
			return
		}

		// To ensure that the token is received in the form `Bearer <token>`
		splitted := strings.Split(token, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Malformed Authorization token.")
			u.BuildResponse(w, 400, response)
		}

		// Extract the token from the splitted token array
		tokenString := splitted[1]
		jwtClaims, err := models.VerifyToken(tokenString)

		if err != nil {
			response = u.Message(false, "Invalid or Expired Authorization token.")
			u.BuildResponse(w, 401, response)
			return
		}

		claims := jwtClaims.(jwt.MapClaims)
		fmt.Println(claims)
		userId := uint(claims["user_id"].(float64))
		// If the token is valid and verified
		ctx := context.WithValue(req.Context(), "user", userId)
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
