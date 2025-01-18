package cmd

import (
	make2 "file_generator/cmd/make"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file_generator",
	Short: "file_generator CLI",
	Long:  "Golang file_generator CLI",
}

func init() {
	rootCmd.AddCommand(
		make2.Make,
	)
}

func Execute() error {
	return rootCmd.Execute()
}
