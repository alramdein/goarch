package cmd

import (
	"fmt"

	"github.com/alramdein/goarch/internal/generator"
	"github.com/spf13/cobra"
)

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize goarch.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing goarch.yaml...")
		generator.GenerateConfig()
	},
}

func init() {
	rootCmd.AddCommand(initConfigCmd)
}
