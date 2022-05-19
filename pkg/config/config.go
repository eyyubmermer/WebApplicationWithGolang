package config

import "html/template"

type AppConfig struct {
	UseCash       bool
	TemplateCashe map[string]*template.Template
}
