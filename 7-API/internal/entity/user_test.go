package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("john Doe", "joj.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "john Doe", user.Name)
	assert.Equal(t, "joj.com", user.Email)
}

func TestUserValidadePassword(t *testing.T) {
	user, err := NewUser("john Doe", "joj.com", "123456")
	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123321"))

	assert.NotEqual(t, user.Password, "123456")
}
