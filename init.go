package main

var commands = map[string]cliCommand{}

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Provides information on commands to use",
			callback:    func(cfg *config) error { return commandHelp() },
		},
		"quit": {
			name:        "quit",
			description: "Ends the pokedex session",
			callback:    func(cfg *config) error { return commandExit() },
		},
		/*
			"map": {
				name:        "map",
				description: "Lists areas",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Go back and lists previous areas",
				callback:    commandMapBack,
			},
		*/
	}
}
