package cmd

import (
	"fmt"

	"github.com/alramdein/goarch/internal/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project structure based on goarch.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating a new project...")
		generator.GenerateProject()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
