// Package banner provides functions for reading ASCII art banners from files.
package banner

import (
	"os"
	"testing"
)

func TestReadBannerFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "valid filename",
			filename: "thinkertoy.txt",
			wantErr:  false,
		},
		{
			name:     "invalid filename",
			filename: "think.txt",
			wantErr:  true,
		},
		{
			name:     "invalid file extension",
			filename: "shadow.png",
			wantErr:  true,
		},
	}
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Chdir("../.."); err != nil {
				t.Fatalf("Failed to change working directory: %v", err)
			}
			// Ensure the working directory is reverted back after the test
			t.Cleanup(func() {
				if err := os.Chdir(originalWD); err != nil {
					t.Fatalf("Failed to revert working directory: %v", err)
				}
			})
			_, err := ReadBannerFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBannerFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
