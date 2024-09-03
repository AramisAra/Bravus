package utils

import (
	"regexp"

	"github.com/AramisAra/BravusServer/config"
)

// This function is the Praser for clientID.
// It uses regexp to parsers the ID to insure
// it's a valid ID.
func IsValidClientString(id string) bool {
	rege := regexp.MustCompile(config.CPattern)
	return rege.MatchString(id)
}

// This function is the Praser for ownerID.
// It uses regexp to parsers the ID to insure
// it's a valid ID.
func IsValidOwnerString(id string) bool {
	rege := regexp.MustCompile(config.OPattern)
	return rege.MatchString(id)
}

// This function is the Praser for serviceID.
// It uses regexp to parsers the ID to insure
// it's a valid ID.
func IsValidServiceString(id string) bool {
	rege := regexp.MustCompile(config.SPattern)
	return rege.MatchString(id)
}
