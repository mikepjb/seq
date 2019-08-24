package seq_test

import (
	"seq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeqCreated(t *testing.T) {
	s := seq.New("e")

	assert.Equal(t, s.ToSlice(), []string{})
}
