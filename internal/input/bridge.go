package input

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//go:embed bin/controller.exe
var controllerBinary []byte

func StartBridge() (func(), error) {
	if runtime.GOOS != "linux" {
		return func() {}, nil
	}

	tmpDir := os.TempDir()
	exePath := filepath.Join(tmpDir, "tetris-controller.exe")

	if err := os.WriteFile(exePath, controllerBinary, 0755); err != nil {
		return nil, fmt.Errorf("write error: %w", err)
	}

	cmd := exec.Command(exePath)

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start error: %w", err)
	}

	return func() {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		os.Remove(exePath)
	}, nil
}
