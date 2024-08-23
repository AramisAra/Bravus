package config

import (
	"os"
)

// This contains the ClientIDPattern
var CPattern = os.Getenv("CIDPattern")

// This contains the OwnerIDPattern
var OPattern = os.Getenv("OIDPattern")

// This contains JWT_Secret use to sign
var Secret = os.Getenv("JWT_SECRET")
