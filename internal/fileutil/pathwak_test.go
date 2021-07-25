package fileutil_test

import (
	"daikincli/internal/fileutil"
	"testing"
)

func TestPathWalk(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				folderPath: "template/test/",
			}, wantErr: true,
		},
		{
			args: args{
				folderPath: "template/workflow/",
			}, wantErr: false, want: "../../template/workflow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileutil.PathWalk(tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathWalk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}
