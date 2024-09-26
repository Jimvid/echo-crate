package middleware

import (
	"context"
	"echo-crate/internal/sessions"
	"encoding/json"
	"log"
	"net/http"

	"github.com/markbates/goth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "auth-session")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userValue := session.Values["user"]
		log.Printf("User value type: %T", userValue)

		var user goth.User
		var err error

		switch v := userValue.(type) {
		case string:
			err = json.Unmarshal([]byte(v), &user)
		case []byte:
			err = json.Unmarshal(v, &user)
		case goth.User:
			user = v
		default:
			log.Printf("Unexpected type for user value: %T", userValue)
			http.Error(w, "User data in unexpected format", http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Printf("Failed to deserialize user data: %v", err)
			http.Error(w, "Failed to retrieve user data", http.StatusInternalServerError)
			return
		}

		log.Printf("User found: %+v", user)

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
