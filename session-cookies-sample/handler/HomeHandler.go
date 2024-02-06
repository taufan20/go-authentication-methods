package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "session-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("Redirect to login page")
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	// Display user's data
	username := session.Values["username"].(string)
	_, _ = fmt.Fprintf(writer, "Welcome, %s!", username)
}
