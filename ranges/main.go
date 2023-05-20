package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// sages := []string{"Gandi", "MLK", "Buddha", "Jesus", "Mohammed"}

	sages := map[string]string{
		"India": "Gandhi",
		"USA": "MLK",
		"India2": "Buddha",
		"Palestine": "Jesus",
		"Arabia": "Mohammed",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}