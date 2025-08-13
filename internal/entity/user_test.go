package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("testuser", "testuser@example.com", "1234567890")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "testuser@example.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("testuser", "testuser@example.com", "1234567890")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("1234567890"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.Equal(t, user.Password, "1234567890")
}
