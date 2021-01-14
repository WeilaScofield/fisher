package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

const text = `
  labels: |-
    app: {{ fromAppEnv "app" }}
    app.kubernetes.io/name: {{ fromAppEnv "name" }}
    app.kubernetes.io/version: {{ fromAppEnv "version" }}
    cafe.sofastack.io/cluster: {{ fromAppEnv "cluster" }}
    cafe.sofastack.io/tenant: {{ fromAppEnv "tenant" }}
    cafe.sofastack.io/workspace: {{ fromAppEnv "workspace" }}
`

var env = map[string]string{
	"app":       "app",
	"namespace": "namespace",
	"name":      "name",
	"version":   "version",
	"cluster":   "cluster",
	"tenant":    "tenant",
	"workspace": "workspace",
}

func main() {
	r, err := renderSidecarTemplate(text, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(r)
}

func renderSidecarTemplate(tmpl string, val interface{}) (string, error) {
	b := &bytes.Buffer{}

	t := template.New("annotations")
	t.Funcs(template.FuncMap{
		"fromAppEnv": fromAppEnv,
	})
	template.Must(t.Parse(tmpl))
	err := t.Execute(b, val)
	return b.String(), err
}

func fromAppEnv(key string) string {
	return env[key]
}
