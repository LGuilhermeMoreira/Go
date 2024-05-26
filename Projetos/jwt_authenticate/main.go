package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func validarCredenciais(name, password string) bool {
	if name == "123321" && password == "123321" {
		return true
	}
	return false
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		// Check for missing authorization header
		if authorizationHeader == "" {
			http.Error(w, "Cabeçalho de autorização ausente", http.StatusUnauthorized)
			return
		}

		// Extract token string from header
		tokenString := strings.Split(authorizationHeader, " ")[1]

		// Parse the JWT token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Replace with your actual secret key stored securely (e.g., environment variables)
			return []byte("sua_chave_secreta"), nil
		})

		// Handle various parsing errors
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				http.Error(w, "Token JWT inválido", http.StatusUnauthorized)
				return
			// case *jwt.ExpiredTokenError:
			// 	http.Error(w, "Token JWT expirado", http.StatusUnauthorized)
			// 	return
			default:
				http.Error(w, "Erro ao processar token JWT", http.StatusInternalServerError)
				return
			}
		}

		// Check if token is valid
		if !token.Valid {
			http.Error(w, "Token JWT inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
			return
		}

		if valido := validarCredenciais(user.Name, user.Password); !valido {
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{
			"userId": user.ID,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte("sua_chave_secreta"))
		if err != nil {
			http.Error(w, "Erro ao gerar token JWT", http.StatusInternalServerError)
			return
		}

		response := map[string]string{"token": signedToken}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	mux.HandleFunc("POST /show", JWTMiddleware(show))

	// Use JWT middleware

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK")
}
