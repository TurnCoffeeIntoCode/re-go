package ssr

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Path            string
	RenderedContent template.HTML
	ClientBundle    template.JS
	Props           map[string]interface{}
}

const pageTemplate = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>React App</title>
	</head>
	<body>
		<div id="app">{{ .RenderedContent }}</div>
		<script type="module">
		{{ .ClientBundle }}
		</script>
		<script>window.PROPS = {{ .Props }};</script>
	</body>
	</html>
`

func convertProps(props map[string]interface{}) (string, error) {
	json, err := json.Marshal(props)
	if err != nil {
		fmt.Println("Error marshalling props")
		return "", errors.New("error marshalling props")
	}
	return string(json), nil
}

func (page *Page) Render(writer http.ResponseWriter) error {
	// Convert props to JSON
	props, err := convertProps(page.Props)
	if err != nil {
		return errors.New("error converting props to react props")
	}
	// Build server side bundle & render the HTML
	html, err := renderHTML(props)
	if err != nil {
		return errors.New("error rendering HTML")
	}

	// Set the rendered content and client bundle
	page.RenderedContent = template.HTML(html)
	page.ClientBundle = template.JS(clientBundleScript)

	// Parse template & render the page
	tmpl, err := template.New("page").Parse(pageTemplate)
	if err != nil {
		fmt.Println(err)
		return errors.New("error parsing template")
	}
	writer.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(writer, page)
	if err != nil {
		fmt.Println(err)
		return errors.New("error executing template")
	}
	return nil
}
