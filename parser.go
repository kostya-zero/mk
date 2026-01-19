package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(strings.NewReader(content))

	var (
		step        Step
		stepName    string
		writingStep bool
		lineNo      int
	)

	for scanner.Scan() {
		lineNo++
		rawLine := scanner.Text()
		line := strings.TrimSpace(rawLine)

		if line == "" {
			if writingStep {
				writingStep = false
				recipe.Steps[stepName] = step
				stepName = ""
				step = Step{}
			}
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if !writingStep {
			left, right, ok := strings.Cut(line, ":")

			if !ok || strings.Contains(right, ":") {
				return nil, &ParseError{
					LineNumber: lineNo,
					Message:    "this is not a valid step initialization syntax",
				}
			}

			step = Step{}
			stepName = strings.TrimSpace(left)
			right = strings.TrimSpace(right)

			if strings.HasSuffix(stepName, "*") {
				step.PassArguments = true
				stepName = strings.TrimSuffix(stepName, "*")
			}

			if right != "" {
				runBefore := strings.Fields(right)
				step.RunBefore = runBefore

				if slices.Contains(runBefore, stepName) {
					return nil, &ParseError{
						LineNumber: lineNo,
						Message:    fmt.Sprintf("possible recursion in step \"%s\" ", stepName),
					}
				}
			}

			writingStep = true
			continue
		}

		step.Commands = append(step.Commands, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, &ParseError{
			LineNumber: lineNo,
			Message:    fmt.Sprintf("scanner failed here: %s", err.Error()),
		}
	}

	// If there is no empty line
	if writingStep {
		recipe.Steps[stepName] = step
	}

	return recipe, nil
}
