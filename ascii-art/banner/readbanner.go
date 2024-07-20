// Package banner provides functions for reading ASCII art banners from files.
package banner

import (
	"io/fs"
	"os"
	"strings"

	"ascii-art-web/ascii-art/errs"
)

// ReadBannerFile reads the ASCII art characters from a banner file and returns them as a slice of strings.
// If the file reading encounters an error, it returns a non-nil error.
func ReadBannerFile(filename string) ([]string, error) {
	filePath := os.DirFS("./ascii-art/banner")
	contentByte, err := fs.ReadFile(filePath, filename)
	if err != nil {
		return nil, err
	}
	if err = errs.CheckFileTamper(filename, contentByte); err != nil {
		return nil, err
	}
	contentString := string(contentByte[1:])
	if filename == "thinkertoy.txt" {
		// convert all carriage returns to newlines
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	return contentSlice, nil
}
