package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
			err := godotenv.Load()

			if tokenString == "" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			if err != nil {
				log.Fatal("Error loading .env file")
			}

			var flag bool
			var hmacSampleSecret = []byte(os.Getenv("JWT_SECRET"))

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					var responseMap = make(map[string]string)
					responseMap["message"] = "Unexpected signing method "
					response, _ := json.Marshal(responseMap)

					w.WriteHeader(http.StatusBadRequest)
					w.Write(response)
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return hmacSampleSecret, nil
			})

			_, ok := token.Claims.(jwt.MapClaims)

			if ok && token.Valid {
				flag = true
			} else {
				flag = false
				fmt.Println(err)
			}

			if flag {
				f(w, r)
			} else {
				var responseMap = make(map[string]string)
				responseMap["message"] = "Invalid JWT"
				response, _ := json.Marshal(responseMap)

				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
		}
	}
}

// Logged request information
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
