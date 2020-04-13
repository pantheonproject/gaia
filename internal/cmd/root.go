package cmd

import (
	"github.com/pantheonproject/gaia/internal/app"
	"github.com/pantheonproject/gaia/internal/dependencies"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var defaultLogger = func() dependencies.Logger { l, _ := zap.NewProduction(); return l }()

type rootCmd struct {
	cmd    *cobra.Command
	logger dependencies.Logger
}

func newRootCmd() (*rootCmd, error) {
	logger, err := zap.NewProduction() // TODO: change logger based on config defaults
	if err != nil {
		return nil, err
	}
	r := &rootCmd{
		logger: logger,
	}
	cmd := &cobra.Command{
		Use:   "gaia",
		Short: "Gaia is a template for Go microservices",
		Long:  "A very simple but functional microservice which serves as a template for other services, with established service wrapping, configuration loading, gRPC/TLS/data-access patterns, etc.",
		Run: func(cmd *cobra.Command, args []string) {
			// Root command logic, for no verbs passed in
			h, err := app.New()
			if err != nil {
				r.getLogger().Error("Failed to create application handle", zap.Error(err))
				return
			}
			err = h.Run()
			if err != nil {
				r.getLogger().Error("Failed to run application", zap.Error(err))
				return
			}
		},
	}
	r.cmd = cmd
	err = r.setupVersionCmd()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *rootCmd) getLogger() dependencies.Logger {
	if r == nil {
		return defaultLogger
	}
	if r.logger == nil {
		r.logger = defaultLogger
	}
	return r.logger
}
