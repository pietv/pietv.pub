package main

import (
	"os"
	"text/template"
)

const ExampleTmpl = `
{{define "Main"}}
AAAAAAAAAAAAAAAA BBBBBBBBBBBBBBBB CCCCCCCCCCCCCCCC
{{end}}
`

func main() {
	template.Must(template.New("Example").
		Parse(ExampleTmpl)).
		ExecuteTemplate(os.Stdout, "Main", nil)
}
