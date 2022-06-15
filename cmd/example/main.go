// team variant 4 (a.k.a. chubaka228)

package main

import (
	"flag"
	"fmt"
	lab2 "github.com/cyberlord-coder-228/SA_lab2"
	"io"
	"log"
	"os"
	"strings"
)

var (
	expression         string
	expressionFilePath string
	outputFilePath     string

	myIn  io.Reader
	myOut io.Writer
)

func initiateFlags() {
	flag.StringVar(&expression, "e", "", "Expression to compute")
	flag.StringVar(&expressionFilePath, "f", "", "File with input")
	flag.StringVar(&outputFilePath, "o", "", "Output file")
	flag.Parse()
}

func main() {
	initiateFlags()

	if expression != "" {
		fmt.Println(expression)
		myIn = strings.NewReader(expression)
	} else if expressionFilePath != "" {
		var openingError error
		myIn, openingError = os.Open(expressionFilePath)
		if openingError != nil {
			log.Fatal(openingError)
			return //openingError
		}
	} else {
		// no input
		// error
		return
	}

	if outputFilePath != "" {
		var creationError error
		myOut, creationError = os.Create(outputFilePath)
		if creationError != nil {
			log.Fatal(creationError)
			return //creationError
		}
	} else {
		myOut = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		InReader:  myIn,
		OutWriter: myOut,
	}

	computationError := handler.Compute()
	if computationError != nil {
		log.Fatal(computationError)
		return
	}

	return
}
