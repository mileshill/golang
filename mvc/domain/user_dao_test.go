package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)
// User Not Found by Id
func TestGetUserNoUserFound(t *testing.T) {
	// Initialization
	// Execution
	user, err := GetUser(0)
	// Validation
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, err.StatusCode, http.StatusNotFound)
}

// User Found by Id
func TestGetUserUserFound(t *testing.T){
	// Initialization
	// Execution
	user, err := GetUser(123)
	// Validation
	assert.NotNil(t, user)
	assert.Nil(t, err)
}
