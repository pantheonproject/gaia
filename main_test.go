package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pantheonproject/gaia/internal/cmd"
)

func Test_main(t *testing.T) {
	ranCount := 0
	startFunc = func() error {
		ranCount++
		return fmt.Errorf("some error")
	}
	defer func() {
		startFunc = cmd.Execute
	}()
	main()
	assert.Equal(t, 1, ranCount)

}
