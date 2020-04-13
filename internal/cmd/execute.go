package cmd

// Execute runs the cobra commander
func Execute() error {
	r, err := newRootCmd()
	if err != nil {
		return err
	}
	return r.cmd.Execute()
}
