package response

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInternalServerError(t *testing.T) {
	assert.Equal(t, http.StatusInternalServerError, InternalServerError("Some error", "").Status)
}

func TestUnauthorized(t *testing.T) {
	assert.Equal(t, http.StatusUnauthorized, Unauthorized("t").Status)
}

func TestInvalidData(t *testing.T) {
	err := InvalidData("Invalid data")
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.NotNil(t, err.Details)
}

func TestNotFound(t *testing.T) {
	assert.Equal(t, http.StatusNotFound, NotFound("abc").Status)
}
