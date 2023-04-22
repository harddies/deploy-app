package model

import "text/template"

type Template struct {
	DeployApp  *template.Template
	ServiceApp *template.Template
}
