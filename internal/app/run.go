package app

// Start creates a new Handle with default config and runs the application on it
func Start() error {
	return New().Run()
}

// Handle stores global state and inter-process comms for an application
type Handle struct {
}

// New creates a new app Handle
func New() *Handle {
	h := &Handle{}
	return h
}

// Run starts the application
func (h *Handle) Run() error {
	return nil
}
