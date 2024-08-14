package main

import (
	"html/template"
	"os"
)

func main() {
	// START1 OMIT
	tmpl := template.Must(template.New("actions").Parse(`
		{{if .IsAdmin}}
			Welcome, Admin {{.Name}}!
		{{else}}
			Welcome, {{.Name}}!
		{{end}}

		Your hobbies:
		{{range .Hobbies}}
			- {{.}}
		{{end}}
	`))
	// END1 OMIT

	// START2 OMIT
	data := struct {
		Name    string
		IsAdmin bool
		Hobbies []string
		Address struct {
			Street string
			City   string
		}
	}{
		Name:    "John",
		IsAdmin: true,
		Hobbies: []string{"gaming", "cooking"},
		Address: struct {
			Street string
			City   string
		}{
			Street: "123 Main St",
			City:   "Goville",
		},
	}
	tmpl.Execute(os.Stdout, data) // HL23
	// END2 OMIT

}
