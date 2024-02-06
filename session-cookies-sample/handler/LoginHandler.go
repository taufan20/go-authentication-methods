package handler

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func LoginHandler(writer http.ResponseWriter, request *http.Request) {

	session, _ := store.Get(request, "session-name")

	// Check if user is already authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}

	// Perform authentication
	username := "user"
	password := "pass"

	if request.FormValue("username") == username &&
		request.FormValue("password") == password {

		session.Values["authenticated"] = true
		session.Values["username"] = username
		err := session.Save(request, writer)
		if err != nil {
			_, _ = fmt.Fprintln(writer, "Unable to save session..")
			return
		}

		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}

	// Display error
	http.Error(writer, "Please enter username and password on x-www-form-urlencoded", http.StatusUnauthorized)

}
