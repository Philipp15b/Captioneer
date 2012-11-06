package ccompiler

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

type CaptionCompiler struct {
	path string
}

func FindCompiler() (*CaptionCompiler, error) {
	path := os.Getenv("sourcesdk")
	if path == "" {
		return nil, errors.New("sourcesdk environment variable not found, make sure the Source SDK is installed!")
	}
	path = path + "/bin/source2007/bin/captioncompiler.exe"
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return &CaptionCompiler{path}, nil
}

func (this CaptionCompiler) Run(target string) (string, error) {
	target, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(this.path, target, "-game", pwd)
	out, err := cmd.Output()
	return string(out), err
}
