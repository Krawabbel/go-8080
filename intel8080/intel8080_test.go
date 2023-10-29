package intel8080_test

import (
	"path/filepath"
	"testing"

	"github.com/Krawabbel/go-8080/cpm"
)

func TestIntel8080(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"non-existent-path", args{"non-existent-path"}, true},
		{"TST8080", args{filepath.Join("..", "cputest", "TST8080.COM")}, false},
		{"8080PRE", args{filepath.Join("..", "cputest", "8080PRE.COM")}, false},
		{"CPUTEST", args{filepath.Join("..", "cputest", "CPUTEST.COM")}, false},
		// {"8080EXM", args{filepath.Join("..", "cputest", "8080EXM.COM")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cpm.Run(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("RunCPM() error = %v,  wantErr %v", err, tt.wantErr)
			}
		})
	}
}
