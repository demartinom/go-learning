package main

import (
	"flag"
	"log"
	"os"

	"link-parser/parser"
)

func main() {
	// Test HTML
	htmlToRead := flag.String("doc", "ex1.html", "html to read")
	flag.Parse()

	// Open HTML file
	file, err := os.Open(*htmlToRead)
	if err != nil {
		log.Fatal(err)
	}
	parser.Parse(file)
}
