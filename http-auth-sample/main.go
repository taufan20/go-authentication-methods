package main

import (
	"encoding/base64"
	"fmt"
	"http-auth-sample/handler"
	"net/http"
	"strings"
)

const AUTHORIZATION = "Authorization"

func basicAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Get authorization header
		authHeader := request.Header.Get(AUTHORIZATION)

		// Check if authorizationheader is not present
		if authHeader == "" {
			writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract encoded base64 user credential by removing prefix "Basic "
		authValue := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(authValue)
		if err != nil {
			http.Error(writer, "Invalid Authorization Header", http.StatusBadRequest)
			return
		}

		// Once decoded, split the string into username and password
		credential := strings.SplitN(string(decoded), ":", 2)
		username := credential[0]
		password := credential[1]

		/* For sample only, the credential check using hardcoded username & password
		* for real project, we can check the credential in database.
		 */
		if !checkCredential(username, password) {
			http.Error(writer, "Credential is Invalid", http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(writer, request)
	})
}

func checkCredential(username string, password string) bool {
	return username == "user" && password == "pass"
}

func main() {
	// Create a new ServerMux
	mux := http.NewServeMux()

	// Attach middleware to the handler
	mux.Handle("/", basicAuthMiddleware(http.HandlerFunc(handler.HelloHandler)))

	fmt.Println("Server is running on: 8080...")
	_ = http.ListenAndServe(":8080", mux)

}
