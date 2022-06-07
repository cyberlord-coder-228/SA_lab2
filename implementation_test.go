// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

func ExampleCalculatePolishNotation() {
	// input must be a valid polish notation expression with operators and operands separated by withespaces
	input := "* 3 + 19 4"

	// function returns result and error
	res, err := CalculatePolishNotation(input)
	
	// Output:
	fmt.Println(res)
	// 69.0
	fmt.Println(err)
	// <nil>
}

func (s *MySuite) TestcCalculatePolishNotation_a(c *C) {
	// 1 + 1
	expression := "+ 1 1"
	result, _ := CalculatePolishNotation(expression)
	expected := 2.0

	c.Assert(result, Equals, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_b(c *C) {
	// 2 ^ 10
	expression := "^ 2 10"
	result, _ := CalculatePolishNotation(expression)
	expected := 1024.0

	c.Assert(result, Equals, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_c(c *C) {
	// 5 + (4 - 2) * 3
	expression := "+ 5 * - 4 2 3"
	result, _ := CalculatePolishNotation(expression)
	expected := 11.0

	c.Assert(result, Equals, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_d(c *C) {
	// 9 + 10 * (3 ^ ((10 - 2) / 4))
	expression := "+ 9 * 10 ^ 3 / - 10 2 4"
	result, _ := CalculatePolishNotation(expression)
	expected := 99.0

	c.Assert(result, Equals, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_e(c *C) {
	// (23 + (98 - 5) * 78) / (11 ^ (7 + 15 - 19))
	expression := "/ + 23 * - 98 5 78 ^ 11 + 7 - 15 19"
	result, _ := CalculatePolishNotation(expression)
	expected := 5.467317806160781

	c.Assert(result, Equals, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_f(c *C) {
	expression := ""
	_, err := CalculatePolishNotation(expression)
	expected := "No input"

	c.Assert(err, ErrorMatches, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_g(c *C) {
	expression := " "
	_, err := CalculatePolishNotation(expression)
	expected := "No input"

	c.Assert(err, ErrorMatches, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_h(c *C) {
	expression := "\n"
	_, err := CalculatePolishNotation(expression)
	expected := "No input"

	c.Assert(err, ErrorMatches, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_i(c *C) {
	expression := "ABOBA"
	_, err := CalculatePolishNotation(expression)
	expected := "Unacceptable character or expression"

	c.Assert(err, ErrorMatches, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_j(c *C) {
	expression := "+ 5 # 3 3"
	_, err := CalculatePolishNotation(expression)
	expected := "Unacceptable character or expression"

	c.Assert(err, ErrorMatches, expected)
}

func (s *MySuite) TestcCalculatePolishNotation_k(c *C) {
	expression := "+- 5 4 3"
	_, err := CalculatePolishNotation(expression)
	expected := "Unacceptable character or expression"

	c.Assert(err, ErrorMatches, expected)
}
