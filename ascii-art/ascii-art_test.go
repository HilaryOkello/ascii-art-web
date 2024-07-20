// Package art returns the ascii art representation
// of a string given a banner filename
package art

import (
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		banner  string
		want    string
		wantErr bool
	}{
		{
			name:   "input with standard",
			str:    "Hello",
			banner: "Standard.txt",
			want: ``,
			wantErr: false,
		},
	}
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Chdir("../."); err != nil {
				t.Fatalf("Failed to change working directory: %v", err)
			}
			// Ensure the working directory is reverted back after the test
			t.Cleanup(func() {
				if err := os.Chdir(originalWD); err != nil {
					t.Fatalf("Failed to revert working directory: %v", err)
				}
			})
			got, err := AsciiArt(tt.str, tt.banner)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsciiArt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AsciiArt() = %v, want %v", len(got), len(tt.want))
			}
		})
	}
}
