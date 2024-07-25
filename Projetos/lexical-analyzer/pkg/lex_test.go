package pkg_test

import (
	"test_lenguage/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLex(t *testing.T) {
	response, err := pkg.Lex("int a = 12;")
	assert.Nil(t, err)
	assert.NotEmpty(t, response)

	response, err = pkg.Lex(`string a = "12"`)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)

	response, err = pkg.Lex(`
		int a = 1;
		int b = 2;
		int c = a + b + 13;
		string A = "abc";
		string B = "def";
		string C = A + B;
	`)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

func TestLexWithError(t *testing.T) {
	response, err := pkg.Lex(`
		int c = a / b;
	`)
	assert.NotNil(t, err)
	assert.Empty(t, response)

	response, err = pkg.Lex(`
		if (c >= 12) {return 12;}
	`)
	assert.NotNil(t, err)
	assert.Empty(t, response)
}
