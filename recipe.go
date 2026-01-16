package main

type Recipe struct {
	Steps map[string]Step
}

type Step struct {
	Commands      []string
	RunBefore     []string
	PassArguments bool
}
