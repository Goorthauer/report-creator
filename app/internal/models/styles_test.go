package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestStyleGetCustomNumStyle(t *testing.T) {
	newStyle := NewFileStyle(excelize.NewFile())

	firstTest := newStyle.GetStyle("datetime")
	assert.Equal(t, 1, firstTest)

	secondTest := newStyle.GetStyle("date")
	assert.Equal(t, 2, secondTest)

	threeTest := newStyle.GetStyle("time")
	assert.Equal(t, 3, threeTest)

	fourTest := newStyle.GetStyle("currency")
	assert.Equal(t, 4, fourTest)

	fiveTest := newStyle.GetStyle("percent")
	assert.Equal(t, 5, fiveTest)

	sixTest := newStyle.GetStyle("")
	assert.Equal(t, 6, sixTest)
}
