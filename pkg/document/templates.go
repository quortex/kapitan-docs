package document

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/quortex/kapitan-docs/pkg/kapitan"
)

// The defaultTemplate to use to render documentation.
const defaultTemplate = `# Classes
{{ range .Classes }}
{{ template "kapitan-docs.displayClass" . }}
{{- end -}}
`

// A define to render a given class.
const defineDisplayClass = `{{/*
Use the default layout to render a class passed as a parameter.
*/}}
{{- define "kapitan-docs.displayClass" -}}
### <a name="{{ .Name }}"></a>{{ .Name }}

{{ escaped "#" .Description }}

{{- with .Uses }}
#### Uses
{{- range . }}
* [{{ . }}](#{{ . }})
{{- end }}
{{- end -}}

{{- with .UsedBy }}
#### Used by
{{- range . }}
* [{{ . }}](#{{ . }})
{{- end }}
{{- end -}}

{{- with .Parameters }}

#### Parameters

| Key | Type | Default | Description |
| --- | ---- | ------- | ----------- |
{{- range . }}
| {{ trimPrefix "parameters." .Key }} | {{ .Kind }} | {{ escaped "|" (multiline .DefaultValue) }} | {{ escaped "|" (multiline .Description) }} |
{{- end }}
{{- end }}
{{- end -}}

`

// tplFunctions is the template functions map to add to the documentation
// template.
func tplFunctions() template.FuncMap {
	return template.FuncMap{
		"multiline":         multiline,
		"escaped":           escaped,
		"classesWithPrefix": classesWithPrefix,
	}
}

// multiline puts a multiline string in preformatted text.
func multiline(s string) string {
	return fmt.Sprintf("<pre>%s</pre>", strings.Replace(s, "\n", "<br />", -1))
}

// escape escape the characters given as parameters in a string.
func escaped(runes string, s string) string {
	toesc := []rune(runes)
	res := ""
	for _, c := range s {
		if contains(toesc, c) {
			res += "\\"
		}
		res += string(c)
	}
	return res
}

// classes With Prefix filters a slice of Class according to a prefix given as a
// parameter.
func classesWithPrefix(classes []kapitan.Class, prefix string) []kapitan.Class {
	res := []kapitan.Class{}
	for _, c := range classes {
		if strings.HasPrefix(c.Name, prefix) {
			res = append(res, c)
		}
	}
	return res
}

// contains returns if given slice contains rune.
func contains(runes []rune, r rune) bool {
	for _, e := range runes {
		if e == r {
			return true
		}
	}
	return false
}
