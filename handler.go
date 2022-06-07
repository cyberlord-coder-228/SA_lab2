// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	"errors"
	"log"
	"io"
	"fmt"
	// "os"
	"strconv"
)

var emptyFileError = errors.New("Empty or non-esixtent imput file")
var noInputError = errors.New("There is no input :/")

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	InReader  io.Reader
	OutWriter io.Writer
}

func (h ComputeHandler) Compute() error {
	buf := make([]byte, 1024)

	n, _ := h.InReader.Read(buf)
	strInput := string(buf[:n])

	fmt.Println(strInput)

	result, calculationError := CalculatePolishNotation(strInput)
	if calculationError != nil {
		log.Fatal(calculationError)
		return calculationError
	}

	strResult := strconv.FormatFloat(result, 'f', -1, 64)

	_, writingError := h.OutWriter.Write([]byte(strResult + "\n"))
	if writingError != nil {
		log.Fatal(writingError)
		return writingError
	}

	return nil
}
