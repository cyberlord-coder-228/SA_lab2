// team variant 4 (a.k.a. chubaka228)

package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var acceptableOperators = []string{"+", "-", "*", "/", "^"}

func reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func contains(slice []string, search string) bool {
	for _, el := range slice {
		if el == search {
			return true
		}
	}
	return false
}

// This function calculates polish notation expresions
func CalculatePolishNotation(input string) (float64, error) {
	if strings.TrimSpace(input) == "" {
		return 0, errors.New("No input")
	}

	var numbersDeck = []float64{}

	symbols := reverse(strings.Fields(input))
	for _, symbol := range symbols {
		if num, err := strconv.ParseFloat(symbol, 64); err == nil {
			// is a number, push it to the deck
			numbersDeck = append(numbersDeck, num)
		} else if contains( /*in*/ acceptableOperators, symbol) {
			// is an operator, do corresponding partial calculation

			// take two numbers from the deck
			deckLen := len(numbersDeck)
			n1 := numbersDeck[deckLen-1]
			n2 := numbersDeck[deckLen-2]
			numbersDeck = numbersDeck[:deckLen-2]

			// calculate
			var res float64
			switch symbol {
			case "+":
				res = n1 + n2
			case "-":
				res = n1 - n2
			case "*":
				res = n1 * n2
			case "/":
				res = n1 / n2
			case "^":
				res = math.Pow(n1, n2)
			}

			// push result to the deck
			numbersDeck = append(numbersDeck, res)
		} else {
			// it`s neither a number nor an acceptable operator sign. hmmmmmmmmm
			return 0, errors.New("Unacceptable character or expression")
		}
	}

	return numbersDeck[0], nil
}
