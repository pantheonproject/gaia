package main

import (
	"fmt"
	"testing"

	"github.com/pantheonproject/gaia/internal/app"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	ranCount := 0
	startFunc = func() error {
		ranCount++
		return fmt.Errorf("some error")
	}
	defer func() {
		startFunc = app.Start
	}()
	main()
	assert.Equal(t, 1, ranCount)

}
