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
		{"Width", "Width", "width", "float64", "float32", GetFunctionTypeNone},
		{"WidthPercent", "WidthPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"Height", "Height", "height", "float64", "float32", GetFunctionTypeNone},
		{"HeightPercent", "HeightPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"PositionType", "PositionType", "position", "PositionType", "flex.PositionType", GetFunctionTypeWithProp},
		{"Direction", "Direction", "direction", "Direction", "flex.Direction", GetFunctionTypeWithProp},
		{"FlexDirection", "FlexDirection", "direction", "FlexDirection", "flex.FlexDirection", GetFunctionTypeWithProp},
		{"JustifyContent", "JustifyContent", "justifyContent", "Justify", "flex.Justify", GetFunctionTypeWithProp},
		{"AlignContent", "AlignContent", "alignContent", "flex.Align", "flex.Align", GetFunctionTypeWithProp},
		{"AlignItems", "AlignItems", "alignItems", "flex.Align", "flex.Align", GetFunctionTypeWithProp},
		{"AlignSelf", "AlignSelf", "alignSelf", "flex.Align", "flex.Align", GetFunctionTypeWithProp},
		{"FlexWrap", "FlexWrap", "wrap", "Wrap", "flex.Wrap", GetFunctionTypeWithProp},
		{"Overflow", "Overflow", "overflow", "Overflow", "flex.Overflow", GetFunctionTypeWithProp},
		{"Display", "Display", "overflow", "Display", "flex.Display", GetFunctionTypeWithProp},
		{"Flex", "Flex", "flex", "float64", "float32", GetFunctionTypeWithProp},
		{"FlexGrow", "FlexGrow", "grow", "float64", "float32", GetFunctionTypeWithGet},
		{"FlexShrink", "FlexShrink", "shrink", "float64", "float32", GetFunctionTypeWithGet},
		{"FlexBasis", "FlexBasis", "basis", "float64", "float32", GetFunctionTypeNone},
		{"FlexBasisPercent", "FlexBasisPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"MinWidth", "MinWidth", "minWidth", "float64", "float32", GetFunctionTypeNone},
		{"MinWidthPercent", "MinWidthPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"MinHeight", "MinHeight", "minHeight", "float64", "float32", GetFunctionTypeNone},
		{"MinHeightPercent", "MinHeightPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"MaxWidth", "MaxWidth", "maxWidth", "float64", "float32", GetFunctionTypeNone},
		{"MaxWidthPercent", "MaxWidthPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"MaxHeight", "MaxHeight", "maxHeight", "float64", "float32", GetFunctionTypeNone},
		{"MaxHeightPercent", "MaxHeightPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"AspectRatio", "AspectRatio", "aspectRatio", "float64", "float32", GetFunctionTypeNone},
	}

	propsWithEdge := []Prop{
		{"Margin", "Margin", "margin", "float64", "float32", GetFunctionTypeWithGet},
		{"MarginPercent", "MarginPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"Padding", "Padding", "padding", "float64", "float32", GetFunctionTypeWithGet},
		{"PaddingPercent", "PaddingPercent", "percent", "float64", "float32", GetFunctionTypeNone},
		{"Border", "Border", "border", "float64", "float32", GetFunctionTypeNone},
		{"Position", "Position", "position", "float64", "float32", GetFunctionTypeWithGet},
		{"PositionPercent", "PositionPercent", "percent", "float64", "float32", GetFunctionTypeNone},
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
