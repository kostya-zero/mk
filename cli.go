package main

import "os"

type Args struct {
	Step    string
	Args    []string
	Help    bool
	Version bool
}

func ParseArgs() (result Args) {
	args := os.Args[1:]
	for index, arg := range args {
		if index == 0 && (arg == "-h" || arg == "--help") {
			result.Help = true
			return
		}

		if index == 0 {
			switch arg {
			case "-h", "--help":
				result.Help = true
				return
			case "-v", "--version":
				result.Version = true
				return
			default:
				result.Step = arg
				continue
			}
		}

		result.Args = append(result.Args, arg)
	}

	if result.Step == "" {
		result.Step = "default"
	}

	return
}
