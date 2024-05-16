package config

import "path/filepath"

const (
	EnvEnvironment = "ENVIRONMENT"
	EnvPort        = "PORT"
	DEVELOPMENT    = "development"
)

var (
	PublicDir = "public"
	PagesDir  = filepath.Join(PublicDir, "pages")
)
