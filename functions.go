package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	// START OMIT
	funcMap := template.FuncMap{
		"toUpper": strings.ToUpper,
		"add":     func(a, b int) int { return a + b },
	}
	tmpl := template.Must(template.New("functions").Funcs(funcMap).Parse(`
		{{.Name | toUpper}}
		{{add 5 3}}
		{{printf "%.2f" .Price}}
	`))
	data := struct {
		Name  string
		Price float64
	}{
		Name:  "pizza",
		Price: 19.99,
	}
	tmpl.Execute(os.Stdout, data) // HL23
	// END OMIT
}
