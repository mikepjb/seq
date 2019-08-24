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
	s := seq.New("E")

	s.Map(strings.ToLower)
	assert.Equal(t, s.Map(strings.ToLower).ToSlice(), []string{"e"})
}
