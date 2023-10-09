/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"epub-ruby-remover/internal"
	"log/slog"
	"path/filepath"

	"github.com/spf13/cobra"
)

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir [directory_path]",
	Short: "Removes ruby tags from all epub files in the specified directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 対象のディレクトリのすべてのepubファイルを取得
		pattern := `*\.epub`
		files, err := filepath.Glob(filepath.Join(args[0], pattern))
		if err != nil {
			slog.Error("failed to get epub files:", err)
		}

		for _, file := range files {
			err := internal.RemoveRuby(file)
			if err != nil {
				slog.Error("failed to remove ruby from epub file:", file, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
}
