package main

import (
	"os"
	"text/template"
)

const ExampleTmpl = `
{{define "Section"}}
{{variable}}
{{end}}

{{define "Main"}}
{{template "Section"}}{{variable}}
{{end}}
`

var helpers = map[string]interface{}{
	"variable": func() string { return "Value" },
}

func main() {
	template.Must(template.New("Example").
		Funcs(helpers).
		Parse(ExampleTmpl)).
		ExecuteTemplate(os.Stdout, "Main", nil)
}
// OMIT
