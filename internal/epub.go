package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

func RemoveRuby(filename string) error {
	// read epub file
	r, err := zip.OpenReader(filename)
	if err != nil {
		return fmt.Errorf("error opening epub file: %w", err)
	}
	defer r.Close()

	// create new epub file
	filenameWithoutExt := strings.TrimSuffix(filename, ".epub")
	newZipFile, err := os.Create(filenameWithoutExt + "_no_ruby.epub")
	if err != nil {
		return fmt.Errorf("error creating new epub file: %w", err)
	}
	defer newZipFile.Close()

	newZipWriter := zip.NewWriter(newZipFile)
	defer newZipWriter.Close()

	slog.Info("Processing EPUB file:", filename)
	slog.Info("Files: ", len(r.File))

	re, err := regexp.Compile(`\.(xhtml|html)$`)
	if err != nil {
		return fmt.Errorf("error compiling regexp: %w", err)
	}

	for i, f := range r.File {
		if match := re.MatchString(f.Name); match {
			slog.Info(fmt.Sprintf("%d / %d - %s", i+1, len(r.File), f.Name))

			rc, err := f.Open()
			if err != nil {
				return fmt.Errorf("error opening file: %w", err)
			}

			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return fmt.Errorf("error reading file: %w", err)
			}

			modifiedData := RemoveRubyTag(string(data))

			// create new file
			newFile, err := newZipWriter.Create(f.Name)
			if err != nil {
				return fmt.Errorf("error creating new file: %w", err)
			}
			_, err = newFile.Write([]byte(modifiedData))
			if err != nil {
				return fmt.Errorf("error writing new file: %w", err)
			}
		} else {
			// Other files are just copied.
			newFile, err := newZipWriter.Create(f.Name)
			if err != nil {
				return fmt.Errorf("error creating new file: %w", err)
			}
			oldFile, err := f.Open()
			if err != nil {
				return fmt.Errorf("error opening file: %w", err)
			}
			_, err = io.Copy(newFile, oldFile)
			oldFile.Close()
			if err != nil {
				return fmt.Errorf("error copying file: %w", err)
			}
		}
	}

	// done
	err = newZipWriter.Close()
	if err != nil {
		return fmt.Errorf("error closing new zip writer: %w", err)
	}
	slog.Info("Done! New EPUB file:", filenameWithoutExt+"_no_ruby.epub")

	return nil
}
