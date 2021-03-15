package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
)

// Holds the application configuration data
type AppConfig struct {
	InfoLog log.Logger
	IsProd	bool
	Session *scs.SessionManager
}
