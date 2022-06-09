// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	// "fmt"
	// "io"
	"os"
	"strings"
	// "testing"
	// "term"

	// . "gopkg.in/check.v1"
)

// func Test(t *testing.T) { TestingT(t) }
// type MySuite struct{}
// var _ = Suite(&MySuite{})

func ExampleCompute() {
	// ComputeHandler.Compute() copies contents from ComputeHandler.Input, calls CalculatePolishNotation(), and writes result to ComputeHandler.Output

	testString := "test"
	testInput := strings.NewReader(testString)
	testOutput := os.Stdout
	testHandler := ComputeHandler{
		InReader:  testInput,
		OutWriter: testOutput,
	}

	testHandler.Compute()

	// Output: "test"
}
