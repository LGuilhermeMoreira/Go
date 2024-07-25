package pkg_test

import (
	"test_lenguage/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidade_Lex(t *testing.T) {
	response, err := pkg.Lex("int a = 12;")
	assert.Nil(t, err)
	assert.NotEmpty(t, response)

	response, err = pkg.Lex(`int a = "12"`)
	assert.NotNil(t, err)
	assert.Empty(t, response)

	response, err = pkg.Lex(`
		int a = 1;
		int b = 2;
		int c = a + b + 13;
		string A = "abc";
		string B = "def";
		string C = A + B; 
	`)
	assert.NotNil(t, err)
	assert.NotEmpty(t, response)

}
