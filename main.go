package main

import (
	"flag"
	"fmt"
	"log"
	"training/gencert/cert"
	"training/gencert/html"
	"training/gencert/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	srcFile := flag.String("file", "", "CSV file input")
	orientation := flag.String("orientation", "landscape", "PDF orientation: portrait|landscape")
	flag.Parse()

	if len(*srcFile) <= 0 {
		log.Fatalf("Invalid file. got=%v\n", *srcFile)
	}

	var saver cert.Saver
	var err error

	switch *outputType {
	case "html":
		if *orientation != "" {
			log.Fatalf("warning: -orientation has no effect for type=%s", *outputType)
		}
		saver, err = html.New("Output")
	case "pdf":
		if *orientation != "lanscape" && *orientation != "portrait" {
			log.Fatalf("warning: -orientation has no effect for %s", *orientation)
		}
		saver, err = pdf.New("Output", *orientation)
	default:
		log.Fatalf("Unknown output type: %s\n", *outputType)
	}

	if err != nil {
		log.Fatalf("Could not create generator: %v\n", err)
	}

	certs, err := cert.ParseCSV(*srcFile)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert, got=%v\n", err)
		}
	}

}
