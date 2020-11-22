package utils

import (
	"fmt"
	"net/url"
)

// GET - HTTP GET Method
const GET string = "GET"

// POST - HTTP POST Method
const POST string = "POST"

// RestrictMethods - Helper for restricting HTTP methods on an endpoint.
func RestrictMethods(allowedMethods []string, method string) error {
	for _, m := range allowedMethods {
		if m == method {
			return nil
		}
	}

	error := fmt.Errorf("%s method not allowed on this endpoint", method)
	return error
}

// GetQueryParamValue - Get the value of a single query param with a matching key
// If no query param matches the provided key, defaultValue is returned
// If multiple matching params exist, defaultValue is returned
func GetQueryParamValue(url *url.URL, key string, defaultValue string) string {
	query := url.Query()
	value, present := query[key]

	if !present || len(value) == 0 {
		return defaultValue
	}

	if len(value) != 1 {
		return defaultValue
	}

	return value[0]
}

// GetQueryParamValues - Get the value of all query params with a matching key
// If no query param matches the provided key, defaultValue is returned
func GetQueryParamValues(url *url.URL, key string, defaultValue []string) []string {
	query := url.Query()
	values, present := query[key]

	if !present || len(values) == 0 {
		return defaultValue
	}

	return values
}
