package handler

import (
	"fmt"
	"net/http"
)

func LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "session-name")

	// Clear the session and redirect to login
	session.Values["authenticated"] = false
	err := session.Save(request, writer)
	if err != nil {
		_, _ = fmt.Fprintln(writer, "Unable to logout session..")
		return
	}

	fmt.Println("Redirect to login page")
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}
