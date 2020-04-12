package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name      string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, Start())
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Handle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New())
		})
	}
}

func TestHandle_Run(t *testing.T) {
	tests := []struct {
		name      string
		h         *Handle
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{}
			tt.assertion(t, h.Run())
		})
	}
}
