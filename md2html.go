package main

import (
	"os"
	"io"
	"io/ioutil"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
)

var (
	w io.Writer = os.Stdout
)

func getMarkdown (filename string) template.HTML {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Could not find markdown file: %s - aborting.", filename)
	}

	contents, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalf("Could not read markdown file: %s - aborting.", filename)
	}

	return template.HTML(github_flavored_markdown.Markdown(contents))
}

func main () {
	template_contents, _ := ioutil.ReadAll(os.Stdin)
	func_map := template.FuncMap{
		"getMarkdown": getMarkdown,
	}

	tmpl, err := template.New("docs").Funcs(func_map).Parse(string(template_contents))

	if err != nil {
		log.Fatalf("Could not parse template. %s", err)
	}

	tmpl.Execute(os.Stdout, nil)
}
