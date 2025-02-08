package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	NextURL string
	PrevURL string
	History []string
}
