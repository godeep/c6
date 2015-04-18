package c6

import "github.com/stretchr/testify/assert"
import "testing"
import "fmt"

func AssertTokenSequence(t *testing.T, l *Lexer, tokenList []TokenType) {
	fmt.Printf("Input: %s\n", l.Input)
	for _, expectingToken := range tokenList {

		fmt.Printf("Expecting token: %s - ", expectingToken.String())

		var token = <-l.Output
		assert.NotNil(t, token)

		if expectingToken == token.Type {
			fmt.Printf("ok (%s)\n", token.Str)
		} else {
			fmt.Printf("not ok (%s)\n", token.Str)
		}
		assert.Equal(t, expectingToken, token.Type)
	}
}