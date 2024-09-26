package sessions

import (
	"encoding/json"
	"net/http"

	"github.com/markbates/goth"
)

func SaveUserToSession(w http.ResponseWriter, r *http.Request, user goth.User) error {
	session, _ := Store.Get(r, "auth-session")

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	session.Values["authenticated"] = true
	session.Values["user"] = string(userData)

	return session.Save(r, w)
}

func GetAuthenticatedUser(r *http.Request) (goth.User, bool) {
	user, ok := r.Context().Value("user").(goth.User)
	return user, ok
}
