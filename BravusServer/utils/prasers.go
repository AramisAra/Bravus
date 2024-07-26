package utils

import (
	"os"
	"regexp"
)

// This Prasers to insure that the ClientId is a valid Format
func IsValidClientString(id string) bool {
	pattern := os.Getenv("ClientIDPattern")
	rege := regexp.MustCompile(pattern)
	return rege.MatchString(id)
}
