package main

import (
	"fmt"
	"os"
)

func runStep(recipe *Recipe, name string) error {
	step, ok := recipe.Steps[name]
	if !ok {
		fmt.Println("step not found")
		os.Exit(1)
	}

	if step.RunBefore != nil {
		for _, stepName := range step.RunBefore {
			err := runStep(recipe, stepName)
			if err != nil {
				return err
			}
		}
	}

	for _, command := range step.Commands {
		err := LaunchCommand(command)
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
	recipe, err := ParseRecipe(dataString)
	if err != nil {
		fmt.Printf("error while parsing recipe: %s\n", err.Error())
		os.Exit(1)
	}

	err = runStep(recipe, args.Step)
	if err != nil {
		fmt.Printf("error while running step '%s': %s\n", args.Step, err.Error())
		os.Exit(1)
	}
}
