package main

var commands = map[string]cliCommand{}

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Provides information on commands to use",
			callback:    commandHelp,
		},
		"quit": {
			name:        "quit",
			description: "Ends the pokedex session",
			callback:    commandExit,
		},

		"map": {
			name:        "map",
			description: "Lists some location areas",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Go back and lists previous areas",
			callback:    commandMapb,
		},
	}
}
