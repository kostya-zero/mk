package main

import (
	"fmt"
	"slices"
	"strings"
)

type ParseError struct {
	LineNumber int
	Message    string
}

func ParseRecipe(content string) (*Recipe, *ParseError) {
	recipe := new(Recipe)
	recipe.Steps = make(map[string]Step)

	// Temporary variable for parsing
	var step Step
	var newStepName string
	var writingStep bool

	for index, line := range strings.Split(content, "\n") {
		trimmedLine := strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		if !writingStep {
			if strings.Contains(trimmedLine, ":") && strings.Count(line, ":") == 1 {
				splittedStep := strings.Split(trimmedLine, ":")
				newStepName = strings.TrimSpace(splittedStep[0])
				leftovers := strings.TrimSpace(splittedStep[1])

				if leftovers != "" {
					runBefore := strings.Split(leftovers, " ")
					step.RunBefore = runBefore
					if slices.Contains(runBefore, newStepName) {
						return nil, &ParseError{
							LineNumber: index + 1,
							Message:    fmt.Sprintf("possible recursion in step \"%s\" ", newStepName),
						}
					}
				}

				if strings.HasSuffix(newStepName, "*") {
					step.PassArguments = true
					newStepName = strings.TrimSuffix(newStepName, "*")
				}

				writingStep = true
			} else {
				return nil, &ParseError{
					LineNumber: index + 1,
					Message:    "this is not a valid step initialization syntax",
				}
			}
		} else {
			if trimmedLine == "" {
				writingStep = false
				recipe.Steps[newStepName] = step
				step = Step{}
				continue
			}
			step.Commands = append(step.Commands, trimmedLine)
		}
	}

	// If there is no empty line
	if writingStep {
		recipe.Steps[newStepName] = step
	}

	return recipe, nil
}
