package brainfuck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrainFuck_NewBrainFuck(t *testing.T) {

	t.Run("should create a brainfuck successfully", func(t *testing.T) {
		brainfuck, err := NewBrainFuck()
		assert.Nil(t, err)
		assert.NotNil(t, brainfuck)
	})

}
