package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"treegen/internal"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var (
	copyToClipboard bool
	outputPath      string
	quiet           bool
)

var rootCmd = &cobra.Command{
	Use:   "treegen <path>",
	Short: "Generate a directory tree structure",
	Long:  "A tool to generate a visual tree representation of a directory.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			return fmt.Errorf("path does not exist: %s", targetPath)
		}

		result := internal.RenderTree(targetPath)

		if copyToClipboard {
			if err := clipboard.WriteAll(result); err != nil {
				return fmt.Errorf("failed to copy to clipboard: %w", err)
			}
			fmt.Fprintln(os.Stderr, "✔  Output copied to clipboard.")
		}

		if outputPath != "" {
			if err := saveToFile(outputPath, result); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "✔  Output saved to: %s\n", outputPath)
		}

		if !quiet {
			fmt.Print(result)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "c", false, "copy the result to the clipboard")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "write the result to a specific `path`")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "suppress stdout output")
}

func saveToFile(path string, content string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
