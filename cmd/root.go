package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "A generator for ",
		Long:  "",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
