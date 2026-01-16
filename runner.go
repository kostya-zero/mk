package main

import (
	"os"
	"os/exec"
	"runtime"
)

type Runner struct {
	shell          string
	launchArgument string
}

func InitRunner() (runner Runner) {
	switch runtime.GOOS {
	case "windows":
		runner.shell = "powershell.exe"
		runner.launchArgument = "-c"
	case "linux":
		runner.shell = "bash"
		runner.launchArgument = "-c"
	case "darwin":
		runner.shell = "zsh"
		runner.launchArgument = "-c"
	default:
		runner.shell = "sh"
		runner.launchArgument = "-c"
	}
	return
}

func (r *Runner) LaunchCommand(args string) error {
	cmd := exec.Command(r.shell, r.launchArgument, args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
