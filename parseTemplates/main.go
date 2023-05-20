package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// Grab files and put them into the  tpl object
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		// Log error
		log.Fatalln(err)
	} 

	nf, err := os.Create("index.html") // write to file
	// err = tpl.Execute(os.Stdout, nil) // write with stdout
	
	if err != nil {
		// Print error
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		// Log error
		log.Fatalln(err)
	} 



}