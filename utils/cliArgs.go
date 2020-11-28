package utils

import (
	"fmt"
)

// IsMigrationEnabled - Checks if "migrate" argument is supplied and returns boolean
func IsMigrationEnabled(args []string) bool {
	return containsArg("migrate", args)
}

// IsFixturesEnabled - Checks if "fixture" argument is supplied and returns boolean
func IsFixturesEnabled(args []string) bool {
	return containsArg("fixture", args)
}

// UseTemporaryDB - Checks if "temp" argument is supplied and returns boolean
func UseTemporaryDB(args []string) bool {
	if containsArg("temp", args) {
		fmt.Println("--- Temporary database will be used ---")
		return true
	}

	return false
}

func containsArg(needle string, haystack []string) bool {
	for _, arg := range haystack {
		if arg == needle {
			return true
		}
	}

	return false
}
