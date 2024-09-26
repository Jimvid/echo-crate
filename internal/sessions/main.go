package sessions

import (
	"github.com/gorilla/sessions"
)

var Store sessions.Store

func init() {
	Store = sessions.NewCookieStore([]byte("your-secret-key")) // Replace with a secure key
}
