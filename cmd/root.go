package cmd

import (
	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use:   "weather",
		Short: "WeatherAPI CLI",
		Long: `WeatherAPI CLi`,
		
	}
)

func Execute() error {
	return rootCmd.Execute()
}
