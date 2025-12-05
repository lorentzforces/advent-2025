package day_01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneSampleInput(t *testing.T) {
	_, err := PartOne("")
	assert.Error(t, err)
}
