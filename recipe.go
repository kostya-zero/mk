package main

type Recipe struct {
	Steps map[string]Step
	Env   map[string]string
}

type Step struct {
	Commands      []string
	RunBefore     []string
	PassArguments bool
}
