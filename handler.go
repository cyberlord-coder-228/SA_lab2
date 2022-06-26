// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	"io"
	"strconv"
)

type ComputeHandler struct {
	InReader  io.Reader
	OutWriter io.Writer
}

func (h ComputeHandler) Compute() error {
	buf := make([]byte, 1024)

	n, _ := h.InReader.Read(buf)
	strInput := string(buf[:n])

	result, calculationError := CalculatePolishNotation(strInput)
	if calculationError != nil {
		return calculationError
	}

	strResult := strconv.FormatFloat(result, 'f', -1, 64)

	_, writingError := h.OutWriter.Write([]byte(strResult))
	if writingError != nil {
		return writingError
	}

	return nil
}
