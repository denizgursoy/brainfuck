package brainfuck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaults_incrementOperation(t *testing.T) {

	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	err := incrementOperation(brainfuck)
	assert.Nil(t, err)
	assert.Equal(t, brainfuck.getCurrentCellValue(), byte(1))
}

func TestDefaults_decrementOperation(t *testing.T) {
	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.Data[brainfuck.DataPointer] = 5

	err := decrementOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.getCurrentCellValue(), byte(4))
}

func TestDefaults_shiftRightOperation(t *testing.T) {
	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4

	err := shiftRightOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.DataPointer, int64(5))
}

func TestDefaults_shiftLeftOperation(t *testing.T) {
	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4

	err := shiftLeftOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.DataPointer, int64(3))
}
