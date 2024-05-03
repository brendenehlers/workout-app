package config

import "path/filepath"

const (
	APP_ENV     = "APP_ENV"
	DEVELOPMENT = "development"
)

var (
	PublicDir = "public"
	PagesDir  = filepath.Join(PublicDir, "pages")
)
