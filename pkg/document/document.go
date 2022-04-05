package document

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/quortex/kapitan-docs/pkg/kapitan"
	log "github.com/sirupsen/logrus"
)

// RenderAsMarkdown  renders given project documentation in markdown format.
func RenderAsMarkdown(p *kapitan.Project, tpl string) (string, error) {
	tpl = stringNotEmptyOr(tpl, defaultTemplate)
	tpl = tpl + defineDisplayClass
	t, err := template.New("template").
		Funcs(sprig.TxtFuncMap()).
		Funcs(tplFunctions()).
		Parse(tpl)

	if err != nil {
		log.Error(err, "Cannot parse template")
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, p); err != nil {
		log.Error(err, "Error parsing generating template output.")
		return "", err
	}

	return buf.String(), nil
}

// stringNotEmptyOr returns first string if not empty or default one.
func stringNotEmptyOr(s string, def string) string {
	if s != "" {
		return s
	}
	return def
}
