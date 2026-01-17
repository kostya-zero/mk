package main

import (
	"fmt"
	"os"
)

func runStep(runner *Runner, recipe *Recipe, name string) error {
	step, ok := recipe.Steps[name]
	if !ok {
		fmt.Println("step not found")
		os.Exit(1)
	}

	if step.RunBefore != nil {
		for _, stepName := range step.RunBefore {
			err := runStep(runner, recipe, stepName)
			if err != nil {
				return err
			}
		}
	}

	for _, command := range step.Commands {
		err := runner.LaunchCommand(command)
		if err != nil {
			return err
		}
	}

	return nil
}

func printHelp() {
	fmt.Println("A task runner.")
	fmt.Println("Usage: mk <FLAG> [TASK] [...ARGUMENTS]\n")
	fmt.Println("Arguments:")
	fmt.Println("    [TASK]            Task to run. Uses default if not provided.")
	fmt.Println("    [ARGUMENTS]       Arguments to pass to the task (if it's allowed).\n")
	fmt.Println("Flags:")
	fmt.Println("    -h, --help        Prints help message")
	fmt.Println("    -v, --version     Prints version")
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

	runner := InitRunner()
	err = runStep(&runner, recipe, cli.Step)
	if err != nil {
		fmt.Printf("error while running step '%s': %s\n", cli.Step, err.Error())
		os.Exit(1)
	}
}
