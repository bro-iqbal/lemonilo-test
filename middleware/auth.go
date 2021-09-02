package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"lemonilo/helpers"

	jwt "github.com/dgrijalva/jwt-go"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Authorization"] != nil {
			authorizationToken := strings.Split(request.Header["Authorization"][0], " ")
			if len(authorizationToken) != 2 {
				helpers.HandleResponse(response, 400, "invalid authorization token", nil)

				return
			}

			if authorizationToken[0] != "Bearer" {
				helpers.HandleResponse(response, 400, "authorization token type does not match", nil)

				return
			}

			var err error
			var token *jwt.Token
			token, err = jwt.Parse(authorizationToken[1], func(token *jwt.Token) (interface{}, error) {
				_, _ = token.Method.(*jwt.SigningMethodHMAC)

				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				helpers.HandleResponse(response, 400, "authorization token credentials do not match", nil)

				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				helpers.HandleResponse(response, 400, "invalid authorization token credentials", nil)

				return
			}

			if token.Valid {
				ctx := context.WithValue(request.Context(), "authorizationToken", claims)

				request = request.WithContext(ctx)

				val := int64(request.Context().Value("authorizationToken").(jwt.MapClaims)["exp"].(float64))

				timeData := time.Unix(val, 0)
				loc, _ := time.LoadLocation("Asia/Jakarta")
				currentTime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), loc)

				if currentTime.After(timeData) {
					helpers.HandleResponse(response, 400, "authorization token has expired", nil)

					return
				}

				next(response, request)
			}
		} else {
			helpers.HandleResponse(response, 401, "unauthorized", nil)

			return
		}
	})
}
