package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestStyleGetCustomNumStyle(t *testing.T) {
	newStyle := NewFileStyle(excelize.NewFile())

	firstTest := newStyle.GetStyle("datetime")
	assert.Equal(t, 1, firstTest)

	secondTest := newStyle.GetStyle("date")
	assert.Equal(t, 2, secondTest)

	ThreeTest := newStyle.GetStyle("time")
	assert.Equal(t, 3, ThreeTest)

	FourTest := newStyle.GetStyle("currency")
	assert.Equal(t, 4, FourTest)

	FiveTest := newStyle.GetStyle("percent")
	assert.Equal(t, 5, FiveTest)

	SixTest := newStyle.GetStyle("")
	assert.Equal(t, 6, SixTest)
}
