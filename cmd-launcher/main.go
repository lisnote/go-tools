package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	exePath, _ := os.Executable()
	currentDir := filepath.Dir(exePath)
	ext := filepath.Ext(exePath)
	cmdPath := strings.TrimSuffix(exePath, ext) + ".cmd"
	cmd := exec.Command("cmd.exe", "/c", cmdPath)
	cmd.Dir = currentDir
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	cmd.Run()
}
