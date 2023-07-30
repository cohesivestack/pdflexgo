package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Prop struct {
	Name         string
	Param        string
	Type         string
	TypeInternal string
}

func main() {

	props := []Prop{
		{"FlexGrow", "grow", "float64", "float32"},
		{"FlexShrink", "shrink", "float64", "float32"},
		{"FlexDirection", "direction", "FlexDirection", "flex.FlexDirection"},
	}

	propsWithEdge := []Prop{
		{"Margin", "margin", "float64", "float32"},
		{"Padding", "padding", "float64", "float32"},
	}

	tmpl, err := template.ParseGlob(filepath.Join("generator", "*.tpl"))
	if err != nil {
		panic(err)
	}

	// Without Edge
	for _, fileName := range []string{
		"block.gen.go",
	} {
		renderTemplate(tmpl, fileName, props)
	}

	// With Edge
	for _, fileName := range []string{
		"block_edge.gen.go",
	} {
		renderTemplate(tmpl, fileName, propsWithEdge)
	}

}

func renderTemplate(tmpl *template.Template, fileName string, accessors []Prop) {
	templateName := fmt.Sprintf("%s.tpl", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(output, templateName, accessors)
	if err != nil {
		panic(err)
	}
}
