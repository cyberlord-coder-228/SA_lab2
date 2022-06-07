// team variant 4 (a.k.a. chubaka228)

package main

import (
	lab2 "https://github.com/cyberlord-coder-228/SA_lab2" // needs to be changed to github link
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var (
	inputExpression    = flag.String("e", "", "Expression to compute")
	expressionFilePath = flag.String("f", "", "File with input")
	outputFilePath     = flag.String("o", "", "Output file")

	myIn  io.Reader
	myOut io.Writer
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	if *inputExpression != "" {
		myIn = strings.NewReader(*inputExpression)
	} else if *expressionFilePath != "" {
		var openingError error
		myIn, openingError = os.Open(*expressionFilePath)
		if openingError != nil {
			log.Fatal(openingError)
			return
		}
	} else {
		// no input
		// error
		return
	}

	if *outputFilePath != "" {
		myOut, openingError := os.Create(*outputFilePath)
		if openingError != nil {
			log.Fatal(openingError)
			return
		}

		defer myOut.Close()
	} else {
		myOut = bufio.NewWriter(os.Stdout)
	}

	handler := &lab2.ComputeHandler{
		InReader:  myIn,
		OutWriter: myOut,
	}

	err := handler.Compute()
}
