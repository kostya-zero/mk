package main

type Args struct {
	Step    string
	Args    []string
	Help    bool
	Version bool
	List    bool
}

func ParseArgs(args []string) Args {
	var result Args

	if len(args) == 0 {
		result.Step = "default"
		return result
	}

	switch args[0] {
	case "-h", "--help":
		result.Help = true
		return result
	case "-v", "--version":
		result.Version = true
		return result
	case "-l", "--list":
		result.List = true
		return result
	default:
		result.Step = args[0]
	}

	if len(args) > 1 {
		result.Args = append(result.Args, args[1:]...)
	}

	if result.Step == "" {
		result.Step = "default"
	}

	return result
}
