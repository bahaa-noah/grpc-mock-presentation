package main

import (
	"html/template"
	"os"
)

// START OMIT
func main() {
	tmpl := template.Must(template.New("example").Parse("Hello, {{.Name}}!"))

	data := struct {
		Name string
	}{
		Name: "Gopher",
	}

	tmpl.Execute(os.Stdout, data) // HL23
}

//END OMIT
