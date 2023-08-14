package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type data struct {
	Name string
	Type string
}

func main() {

	tmpl, err := template.ParseGlob(filepath.Join("generator", "*.tpl"))
	if err != nil {
		panic(err)
	}

	for _, name := range []string{
		"block",
		"page",
	} {
		renderTemplate(tmpl, "container", name)
	}

	for _, name := range []string{
		"block",
		"text",
		"text_multi_format",
		"image",
	} {
		renderTemplate(tmpl, "node", name)
	}

	for _, name := range []string{
		"page",
		"block",
		"text",
		"text_multi_format",
		"image",
	} {
		renderTemplate(tmpl, "element", name)
	}
}

func renderTemplate(tmpl *template.Template, group string, name string) {
	templateName := fmt.Sprintf("%s.gen.go.tpl", group)

	output, err := os.Create(fmt.Sprintf("%s_%s.gen.go", name, group))
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(
		output,
		templateName,
		data{
			Name: snakeToCamel(name),
			Type: fmt.Sprintf("%s%s", upperFirst(snakeToCamel(name)), "Element")})
	if err != nil {
		panic(err)
	}
}

func upperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func snakeToCamel(s string) string {
	words := strings.Split(s, "_")

	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			r := []rune(words[i])
			r[0] = unicode.ToTitle(r[0])
			words[i] = string(r)
		}
	}

	return strings.Join(words, "")
}
