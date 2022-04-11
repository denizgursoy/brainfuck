package brainfuck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
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
	t.Run("should shift to right if position is not at last", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		brainfuck.Data = append(brainfuck.Data, 1, 2, 3, 4)

		brainfuck.DataPointer = 0

		err := shiftRightOperation(brainfuck)

		assert.Nil(t, err)
		assert.Equal(t, brainfuck.DataPointer, int64(1))
	})

	t.Run("should increase size of data slice if position is at last", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		brainfuck.DataPointer = InitialCapacity - 1

		err := shiftRightOperation(brainfuck)

		assert.Nil(t, err)
		assert.Equal(t, brainfuck.DataPointer, int64(InitialCapacity))
		assert.Equal(t, len(brainfuck.Data), InitialCapacity+1)
	})
}

func TestDefaults_shiftLeftOperation(t *testing.T) {
	t.Run("should shift to left if the current poisiton is not 0", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		brainfuck.DataPointer = 4

		err := shiftLeftOperation(brainfuck)

		assert.Nil(t, err)
		assert.Equal(t, brainfuck.DataPointer, int64(3))
	})

	t.Run("should return error if the current position is 0", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		brainfuck.DataPointer = 0

		err := shiftLeftOperation(brainfuck)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ShiftLeftNoSpaceError)
	})
}

func TestDefaults_printOperation(t *testing.T) {
	ioOptions := createIoOptions()

	buffer := bytes.Buffer{}
	ioOptions.OutputWriter = &buffer

	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4
	brainfuck.Data[brainfuck.DataPointer] = 16

	err := printOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, buffer.Bytes()[0], byte(16))
}

func TestDefaults_setFromUserInputOperation(t *testing.T) {
	ioOptions := createIoOptions()

	ioOptions.InputReader = strings.NewReader("input")

	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4
	brainfuck.Data[brainfuck.DataPointer] = 16

	err := setFromUserInputOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.getCurrentCellValue(), byte(105))
}

func TestDefaults_startLoopOperation(t *testing.T) {
	t.Run("should go to the end of loop if the cell value is 0", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		start := int64(5)
		end := int64(15)

		brainfuck.loopStack = append(brainfuck.loopStack, &Loop{
			Start: &start,
			End:   &end,
		})

		brainfuck.CommandPointer = start

		brainfuck.DataPointer = 4
		brainfuck.Data[brainfuck.DataPointer] = 0

		err := startLoopOperation(brainfuck)

		assert.Nil(t, err)
		assert.Equal(t, brainfuck.CommandPointer, end)
	})

	t.Run("should add new loop to stack if loop is nested",
		func(t *testing.T) {
			ioOptions := createIoOptions()
			brainfuck, _ := NewBrainFuck(ioOptions)

			start := int64(5)
			end := int64(15)

			brainfuck.DataPointer = 4
			brainfuck.Data[brainfuck.DataPointer] = 12

			brainfuck.loopStack = append(brainfuck.loopStack, &Loop{
				Start: &start,
				End:   &end,
			})

			brainfuck.CommandPointer = int64(10)

			err := startLoopOperation(brainfuck)

			assert.Nil(t, err)
			assert.Equal(t, len(brainfuck.loopStack), 2)
			assert.Equal(t, *brainfuck.loopStack[1].Start, brainfuck.CommandPointer)
			assert.Nil(t, brainfuck.loopStack[1].End)
		})

	t.Run("should add loop if the stack is empty", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		brainfuck.CommandPointer = int64(10)

		brainfuck.DataPointer = 4
		brainfuck.Data[brainfuck.DataPointer] = 12

		err := startLoopOperation(brainfuck)

		assert.Nil(t, err)
		assert.Equal(t, len(brainfuck.loopStack), 1)
		assert.Equal(t, *brainfuck.loopStack[0].Start, brainfuck.CommandPointer)
		assert.Nil(t, brainfuck.loopStack[0].End)
	})
}

func TestDefaults_endLoopOperation(t *testing.T) {

	t.Run("should set end to current command pointer", func(t *testing.T) {

		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		start := int64(5)

		brainfuck.DataPointer = 4
		brainfuck.Data[brainfuck.DataPointer] = 12

		brainfuck.loopStack = append(brainfuck.loopStack, &Loop{
			Start: &start,
			End:   nil,
		})

		currentCommandPointer := int64(10)
		brainfuck.CommandPointer = currentCommandPointer

		err := endLoopOperation(brainfuck)
		assert.Nil(t, err)
		assert.Equal(t, *brainfuck.loopStack[0].End, currentCommandPointer)

	})

	t.Run("should return to the beginning of loop if value is non zero", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		start := int64(5)
		end := int64(15)

		brainfuck.DataPointer = 4
		brainfuck.Data[brainfuck.DataPointer] = 12

		brainfuck.loopStack = append(brainfuck.loopStack, &Loop{
			Start: &start,
			End:   &end,
		})

		err := endLoopOperation(brainfuck)
		assert.Nil(t, err)
		assert.Equal(t, brainfuck.CommandPointer, start)
	})

	t.Run("should pop from stack if the value is 0, and pointer should not change", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		start := int64(5)
		end := int64(15)

		brainfuck.DataPointer = 4
		brainfuck.Data[brainfuck.DataPointer] = 0

		brainfuck.loopStack = append(brainfuck.loopStack, &Loop{
			Start: &start,
			End:   &end,
		})

		brainfuck.CommandPointer = end

		err := endLoopOperation(brainfuck)
		assert.Nil(t, err)
		assert.Equal(t, brainfuck.CommandPointer, end)
		assert.Equal(t, len(brainfuck.loopStack), 0)
	})

	t.Run("should return error is stack is empty", func(t *testing.T) {
		ioOptions := createIoOptions()
		brainfuck, _ := NewBrainFuck(ioOptions)

		brainfuck.DataPointer = 4
		brainfuck.Data[brainfuck.DataPointer] = 0

		err := endLoopOperation(brainfuck)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, LoopEndInvalidError)
	})
}
