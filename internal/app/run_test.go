package app

import (
	"testing"

	"github.com/pantheonproject/gaia/internal/dependencies"
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
		name      string
		want      *Handle
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHandle_Run(t *testing.T) {
	type fields struct {
		logger dependencies.Logger
	}
	tests := []struct {
		name      string
		fields    fields
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{
				logger: tt.fields.logger,
			}
			tt.assertion(t, h.Run())
		})
	}
}

func TestHandle_getLogger(t *testing.T) {
	type fields struct {
		logger dependencies.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   dependencies.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{
				logger: tt.fields.logger,
			}
			assert.Equal(t, tt.want, h.getLogger())
		})
	}
}
