package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(writer, "Hello! your account has been granted to the resource. Thank You.")
}
