// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var emptyFileError = errors.New("Empty or non-esixtent imput file")
var noInputError = errors.New("There is no input :/")

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	InReader  string
	OutWriter string
}

func (h ComputeHandler) Compute() error {
	input, _ := InReader.Read()

	result, calculationError := CalculatePolishNotation(input)
	if calculationError != nil {
		log.Fatal(calculationError)
		return calculationError
	}

	// strResult := strconv.FormatFloat(result, 'f', -1, 64)

	_, writingError := OutWriter.Write([]byte(result))
	if writingError != nil {
		log.Fatal(writingError)
		return writingError
	}

	return nil
}
