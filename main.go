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

func main() {
	args := ParseArgs()

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
	err = runStep(&runner, recipe, args.Step)
	if err != nil {
		fmt.Printf("error while running step '%s': %s\n", args.Step, err.Error())
		os.Exit(1)
	}
}
