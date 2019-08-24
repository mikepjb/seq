package seq_test

import (
	"seq"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeqCreated(t *testing.T) {
	assert.Equal(t, seq.New("e").ToSlice(), []string{"e"})
}

func TestMap(t *testing.T) {
	s := seq.New("East", "West")
	s.Map(strings.ToLower)
	assert.Equal(t, s.Map(strings.ToLower).ToSlice(), []string{"east", "west"})
}

func TestMapChain(t *testing.T) {
	s := seq.New("East London")
	result := s.Map(strings.ToLower).
		Map(strings.Replace, " ", "-", -1).
		ToSlice()

	assert.Equal(t, []string{"east-london"}, result)
}
