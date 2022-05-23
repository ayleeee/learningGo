package config

import (
	"html/template"
	"log"
)

//imported by other application, not importing something

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
