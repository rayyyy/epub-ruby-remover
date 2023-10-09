/*
Copyright © 2023 Ray Navarro nraavy@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"archive/zip"
	"epub-ruby-remover/internal"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "epub-ruby-remover [epub_file]",
	Short: "Removes ruby tags from epub files",
	Long:  `Removes ruby tags from epub files.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		// 拡張子を取得してepubかどうかを判定する
		if match, _ := regexp.MatchString(`\.epub$`, args[0]); !match {
			return fmt.Errorf("invalid file extension")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// read epub file
		r, err := zip.OpenReader(args[0])
		if err != nil {
			slog.Error("Error opening EPUB file:", err)
			return
		}
		defer r.Close()

		// create new epub file
		filenameWithoutExt := strings.TrimSuffix(args[0], ".epub")
		newZipFile, err := os.Create(filenameWithoutExt + "_no_ruby.epub")
		if err != nil {
			slog.Error("Error creating new EPUB file:", err)
			return
		}
		defer newZipFile.Close()

		newZipWriter := zip.NewWriter(newZipFile)
		defer newZipWriter.Close()

		slog.Info("Processing EPUB file:", args[0])
		slog.Info("Files: ", len(r.File))

		re, err := regexp.Compile(`\.(xhtml|html)$`)
		if err != nil {
			slog.Error("Error compiling regexp:", err)
			return
		}

		for i, f := range r.File {
			if match := re.MatchString(f.Name); match {
				slog.Info(fmt.Sprintf("%d / %d - %s", i+1, len(r.File), f.Name))

				rc, err := f.Open()
				if err != nil {
					slog.Error("Error opening file:", err)
					return
				}

				data, err := io.ReadAll(rc)
				rc.Close()
				if err != nil {
					slog.Error("Error reading file:", err)
					return
				}

				modifiedData := internal.RemoveRuby(string(data))

				// 新しいZIPファイルに書き込む
				newFile, err := newZipWriter.Create(f.Name)
				if err != nil {
					slog.Error("Error creating new file:", err)
					return
				}
				_, err = newFile.Write([]byte(modifiedData))
				if err != nil {
					slog.Error("Error writing file:", err)
					return
				}
			} else {
				// XHTML/HTML以外のファイルはそのままコピー
				newFile, err := newZipWriter.Create(f.Name)
				if err != nil {
					slog.Error("Error creating new file:", err)
					return
				}
				oldFile, err := f.Open()
				if err != nil {
					slog.Error("Error opening file:", err)
					return
				}
				_, err = io.Copy(newFile, oldFile)
				oldFile.Close()
				if err != nil {
					slog.Error("Error copying file:", err)
					return
				}
			}
		}

		// done
		err = newZipWriter.Close()
		if err != nil {
			slog.Error("Error closing new EPUB file:", err)
			return
		}
		slog.Info("Done! New EPUB file:", filenameWithoutExt+"_no_ruby.epub")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
