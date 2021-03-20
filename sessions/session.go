package sessions

import (
	"crypto/rand"
	"io"
	"net/http"
	"github.com/gorilla/sessions"
)


var Store = sessions.NewCookieStore(generateRandomCookieKey(64))

func generateRandomCookieKey(length int) []byte {
	maker := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, maker); err != nil {
		return nil
	}
	return maker
}

func SetSession(w http.ResponseWriter, r *http.Request, username string) error {
	session, _ := Store.Get(r, "session")
	session.Values["username"] = username
	return session.Save(r, w)
}