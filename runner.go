package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Runner struct {
	shell          string
	launchArgument string
	env            []string
}

func InitRunner(env map[string]string) (runner Runner) {
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

	if len(env) > 0 {
		for key, value := range env {
			runner.env = append(runner.env, fmt.Sprintf("%s=%s", key, value))
		}
	}

	return
}

func (r *Runner) LaunchCommand(args string) error {
	cmd := exec.Command(r.shell, r.launchArgument, args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), r.env...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
