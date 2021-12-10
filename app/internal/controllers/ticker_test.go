package controllers

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//Test_removeFilesFromDirectory fictive test
func Test_removeFilesFromDirectory(t *testing.T) {
	removeFilesFromDirectory(time.Now())
}

func Test_removeFIles(t *testing.T) {
	for i := 0; i < 50; i++ {
		testName := fmt.Sprintf("test_%v_name_.xlsx", i)
		f, err := os.Create(testName)
		assert.Nil(t, err)
		err = f.Close()
		assert.Nil(t, err)
		_, err = removeFiles(testName)
		assert.Nil(t, err)
	}
}
