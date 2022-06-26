// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	"os"
	"strings"
	"bytes"

	. "gopkg.in/check.v1"
)

func ExampleComputeHandler_Compute() {
	// ComputeHandler.Compute() copies contents from ComputeHandler.Input, calls CalculatePolishNotation(), and writes result to ComputeHandler.Output

	testHandler := ComputeHandler{
		InReader: strings.NewReader("* 3 23"),
		OutWriter: os.Stdout,
	}

	testHandler.Compute()

	// Output: 69
}

func (s *MySuite) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		InReader:  strings.NewReader("+ 2 2"),
		OutWriter: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "4")
}

func (s *MySuite) TestComputeHandlerError1(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		InReader:  strings.NewReader("23 22"),
		OutWriter: b,
	}
	err := handler.Compute()

	c.Assert(err, ErrorMatches, "More final values than expected")
}

func (s *MySuite) TestComputeHandlerError2(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		InReader:  strings.NewReader(""),
		OutWriter: b,
	}
	err := handler.Compute()

	c.Assert(err, ErrorMatches, "No input")
}

func (s *MySuite) TestComputeHandlerError3(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		InReader:  strings.NewReader("& 1 4"),
		OutWriter: b,
	}
	err := handler.Compute()

	c.Assert(err, ErrorMatches, "Unacceptable character or expression")
}

func (s *MySuite) TestComputeHandlerError4(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		InReader:  strings.NewReader("& a 4"),
		OutWriter: b,
	}
	err := handler.Compute()

	c.Assert(err, ErrorMatches, "Unacceptable character or expression")
}

func (s *MySuite) TestComputeHandlerError5(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		InReader:  strings.NewReader("qwerty qwe rty"),
		OutWriter: b,
	}
	err := handler.Compute()

	c.Assert(err, ErrorMatches, "Unacceptable character or expression")
}
