package main

import (
	"html/template"
	"os"
)

func main() {
	// START OMIT
	tmpl := template.Must(template.New("basic").Parse(`
		Name: {{.Name}}
		Age: {{.Age}}
		Hobbies: {{.Hobbies}}
	`))
	data := struct {
		Name    string
		Age     int
		Hobbies []string
	}{
		Name:    "Alice",
		Age:     21,
		Hobbies: []string{"reading", "coding", "hiking"},
	}
	tmpl.Execute(os.Stdout, data) // HL23
	// END OMIT
}
