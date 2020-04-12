package app

import (
	"github.com/pantheonproject/gaia/internal/dependencies"
	"go.uber.org/zap"
)

var defaultLogger = func() dependencies.Logger { l, _ := zap.NewProduction(); return l }()

// Start creates a new Handle with default config and runs the application on it
func Start() error {
	h, err := New()
	if err != nil {
		defaultLogger.Error("Failed to start", zap.Error(err))
		return err
	}
	err = h.Run()
	if err != nil {
		defaultLogger.Error("Failed to run", zap.Error(err))
		return err
	}
	return nil
}

// Handle stores global state and inter-process comms for an application
type Handle struct {
	logger dependencies.Logger
}

// New creates a new app Handle
func New() (*Handle, error) {
	h := &Handle{}
	logger, err := zap.NewProduction()
	if logger != nil {
		h.logger = logger
	}
	return h, err
}

// Run starts the application
func (h *Handle) Run() error {
	h.getLogger().Info("Hello, world!")
	return nil
}

func (h *Handle) getLogger() dependencies.Logger {
	if h == nil {
		return defaultLogger
	}
	if h.logger == nil {
		h.logger = defaultLogger
	}
	return h.logger
}
