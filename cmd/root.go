package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "golang project template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Greetings!")
	},
}

// Execute execcutes commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
