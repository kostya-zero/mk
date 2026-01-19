package main

import (
	"fmt"
	"os"
	"strings"
)

func runStep(runner *Runner, recipe *Recipe, name string, args string) error {
	step, ok := recipe.Steps[name]
	if !ok {
		fmt.Println("step not found")
		os.Exit(1)
	}

	if step.RunBefore != nil {
		for _, stepName := range step.RunBefore {
			err := runStep(runner, recipe, stepName, args)
			if err != nil {
				return err
			}
		}
	}

	for _, command := range step.Commands {
		command = strings.ReplaceAll(command, "$*", args)
		err := runner.LaunchCommand(command)
		if err != nil {
			return err
		}
	}

	return nil
}

func printHelp() {
	fmt.Println("A task runner.")
	fmt.Printf("Usage: mk <FLAG> [TASK] [...ARGUMENTS]\n\n")
	fmt.Println("Arguments:")
	fmt.Println("    [TASK]            Task to run. Uses default if not provided.")
	fmt.Printf("    [ARGUMENTS]       Arguments to pass to the task (if it's allowed).\n\n")
	fmt.Println("Flags:")
	fmt.Println("    -h, --help        Prints help message")
	fmt.Println("    -v, --version     Prints version")
	fmt.Println("    -l, --list        Prints version")
}

func main() {
	args := os.Args
	cli := ParseArgs(args[1:])

	if cli.Help {
		printHelp()
		return
	}

	if cli.Version {
		fmt.Printf("mk %s", version)
		return
	}

	data, err := os.ReadFile("Mkfile")
	if err != nil {
		fmt.Printf("failed to read mkfile: %s\n", err.Error())
		os.Exit(1)
	}

	dataString := string(data)
	recipe, parseError := ParseRecipe(dataString)
	if parseError != nil {
		fmt.Printf("[Mkfile:%d] error: %s\n", parseError.LineNumber, parseError.Message)
		os.Exit(1)
	}

	if cli.List {
		for step := range recipe.Steps {
			fmt.Println(step)
		}
		return
	}

	if cli.Env {
		for key, value := range recipe.Env {
			fmt.Printf("%s=%s\n", key, value)
		}
		return
	}

	var stepArgs string
	for _, a := range cli.Args {
		stepArgs = stepArgs + a + " "
	}
	stepArgs = strings.TrimSpace(stepArgs)
	runner := InitRunner(recipe.Env)
	err = runStep(&runner, recipe, cli.Step, stepArgs)
	if err != nil {
		fmt.Printf("error while running step '%s': %s\n", cli.Step, err.Error())
		os.Exit(1)
	}
}
