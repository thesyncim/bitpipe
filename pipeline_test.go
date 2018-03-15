package bitpipe

import (
	"os"
	"reflect"
	"testing"
)

func TestPipeline_Clone(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		{
			name: "clone revision",
			p: &Pipeline{
				OutputStream: os.Stdout,
				RepoURL:      "https://github.com/src-d/go-git.git",
				Revision:     "e39559fba6ea936c5f4659c03c22f9a5a256a4d7",
			},
			wantErr: false,
		},
		// TODO:improve test
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.clone(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.Clone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline_Run(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline_pullImage(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.pullImage(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.pullImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline_createContainer(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.createContainer(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.createContainer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline_startContainer(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.startContainer(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.startContainer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline_stopContainer(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.stopContainer(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.stopContainer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipeline_removeContainer(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pipeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.removeContainer(); (err != nil) != tt.wantErr {
				t.Errorf("Pipeline.removeContainer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_img2RepoandTag(t *testing.T) {
	type args struct {
		img string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := img2RepoandTag(tt.args.img)
			if got != tt.want {
				t.Errorf("img2RepoandTag() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("img2RepoandTag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_loadEnvFile(t *testing.T) {
	type args struct {
		envfile string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadEnvFile(tt.args.envfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadEnvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadEnvFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
