// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import (
	"github.com/kjk/flex"
)

{{ range . -}}
func (block *BlockElement) {{ .Name }}({{ if and .Param .Type }}{{ .Param }} {{ .Type }}{{ end }}) *BlockElement {
	block.getFlexNode().StyleSet{{ .NameInternal }}({{ if and .Param .TypeInternal }}{{ .TypeInternal }}({{ .Param }}){{ end }})
	return block
}
{{if or (eq .GetFunctionType "withGet") (eq .GetFunctionType "withProp")}}

func (block *BlockElement) Get{{ .Name }}() {{ .Type }} {
	return {{ .Type }}(block.getFlexNode().{{ if eq .GetFunctionType "withGet" }}StyleGet{{ .NameInternal }}(){{ else }}Style.{{ .NameInternal }}{{ end }})
}
{{ end }}

{{ end }}