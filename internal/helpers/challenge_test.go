package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	testChallenge := []byte("fba5113f-71cc-4b50-8354-f4731b1a3b6e")
	expectedPrefix := []byte("200362")
	actualPrefix := Solve(testChallenge)
	assert.Equal(t, expectedPrefix, actualPrefix)
}

func TestVerify(t *testing.T) {
	testChallenge := []byte("fba5113f-71cc-4b50-8354-f4731b1a3b6e")
	testPrefix := []byte("200362")
	assert.True(t, Verify(testChallenge, testPrefix))
}
