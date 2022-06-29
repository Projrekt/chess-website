package data

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"go-projects/chess/service"
	"go-projects/chess/util"

	uuid "github.com/satori/go.uuid"
)

// Encrypt a password
func Encrypt(text string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(text)))
	return
}

// CreateUUID to store in a cookie
func CreateUUID() string {
	sID := uuid.NewV4()
	return sID.String()
}

// AssignCookie puts a cookie into the response writer using the session uuid as the value
func AssignCookie(w http.ResponseWriter, r *http.Request, sess service.Session) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    sess.Uuid,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

func DeleteSession(w http.ResponseWriter, r *http.Request, serve *service.DbService) {
	cookie, err := r.Cookie("session")
	util.ErrHandler(err, "DeleteSession", "Session", time.Now(), w)

	// remove cookie from the browser
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	// remove the session from the database
	session := service.Session{Uuid: cookie.Value}
	if serve != nil {
		serve.DeleteByUUID(session)
	}
	http.Redirect(w, r, "/", 302)

}

// AuthSession checks if a users password matches the password for the user in the db
// then creates a session and sets the cookie in the browser
func AuthSession(w http.ResponseWriter, r *http.Request, u service.User, serve *service.DbService) {
	if u.Password == Encrypt(r.PostFormValue("password")) {
		session, err := serve.CreateSession(u)
		util.ErrHandler(err, "CreateSession", "Database", time.Now(), w)
		AssignCookie(w, r, session)
	} else {
		err := fmt.Errorf("Bad password")
		util.ErrHandler(err, "Authenticate", "Password", time.Now(), w)
	}
}
