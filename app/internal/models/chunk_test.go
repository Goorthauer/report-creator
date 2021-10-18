package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkTypeAssignValue(t *testing.T) {
	chunk := NewChunkType()

	chunk.AssignValue(7)
	assert.Equal(t, 7, chunk.GetValue())

	chunk.AssignValue(5)
	assert.Equal(t, 5, chunk.GetValue())

}

func TestChunkType_GetCountPages(t *testing.T) {

	chunk := NewChunkType()
	chunkInt, err := chunk.GetCountPages(5)

	assert.Equal(t, 1, chunkInt)

	assert.Nil(t, err)
}
