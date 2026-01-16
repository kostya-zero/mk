package main

import (
	"os"
	"os/exec"
	"runtime"
)

func LaunchCommand(args string) error {
	var shell string
	var launchArgument string

	switch runtime.GOOS {
	case "windows":
		shell = "cmd.exe"
		launchArgument = "/C"
	case "linux":
		shell = "bash"
		launchArgument = "-c"
	case "darwin":
		shell = "zsh"
		launchArgument = "-c"
	default:
		shell = "sh"
		launchArgument = "-c"
	}

	cmd := exec.Command(shell, launchArgument, args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
