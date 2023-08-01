package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type GetFunctionType string

const (
	GetFunctionTypeNone     GetFunctionType = "none"
	GetFunctionTypeWithGet  GetFunctionType = "withGet"
	GetFunctionTypeWithProp GetFunctionType = "withProp"
)

type Prop struct {
	Name            string
	NameInternal    string
	Param           string
	Type            string
	TypeInternal    string
	GetFunctionType GetFunctionType
}

func main() {

	props := []Prop{
		{"FlexDirection", "FlexDirection", "direction", "FlexDirection", "flex.FlexDirection", GetFunctionTypeWithProp},
		{"FlexWrap", "FlexWrap", "wrap", "Wrap", "flex.Wrap", GetFunctionTypeWithProp},
		{"Flex", "Flex", "flex", "float64", "float32", GetFunctionTypeWithProp},
		{"FlexGrow", "FlexGrow", "grow", "float64", "float32", GetFunctionTypeWithGet},
		{"FlexShrink", "FlexShrink", "shrink", "float64", "float32", GetFunctionTypeWithGet},
		{"FlexBasis", "FlexBasis", "basis", "float64", "float32", GetFunctionTypeNone},
		{"FlexBasisPercent", "FlexBasisPercent", "basisPercent", "float64", "float32", GetFunctionTypeNone},
	}

	propsWithEdge := []Prop{
		{"Margin", "Margin", "margin", "float64", "float32", GetFunctionTypeWithGet},
		{"Padding", "Padding", "padding", "float64", "float32", GetFunctionTypeWithGet},
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
