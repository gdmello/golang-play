package main

import (
	"github.com/Masterminds/sprig"
	"html/template"
	"bytes"
)

/* 
Based on http://technosophos.com/2016/03/29/go-quickly-50-template-functions-from-sprig.html

*/

func main() {
	// The template
	tpl := " Heavenly {{.Name}} image: {{regexSplit \":\" .Iag -1 | last }}"
	
	// The values to pass to the template
	vars := map[string]string{"Name": "Home", "Iag": "registry.points.com/prod/platform_core:2018-23-12"}
	
	// Compile the template
	// t := template.Must(template.New("example").Parse(tpl))
	t := template.Must(template.New("too").Funcs(sprig.FuncMap()).Parse(tpl))

	// Run the template, copying the output to the buffer.
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
			panic(err)
	}
	print(b.String())
}