package utils

import (
	"fmt"
	"net/http"
)

// GET - HTTP GET Method
const GET string = "GET"

// POST - HTTP POST Method
const POST string = "POST"

// RestrictMethods - Helper for restricting HTTP methods on an endpoint.
// Will write a 405 response if incorrect method is requested
func RestrictMethods(allowedMethods []string, method string, w http.ResponseWriter) error {
	for _, m := range allowedMethods {
		if m == method {
			return nil
		}
	}

	error := fmt.Errorf("%s method not allowed on this endpoint", method)
	w.WriteHeader(405)
	fmt.Fprintf(w, error.Error())
	return error
}
