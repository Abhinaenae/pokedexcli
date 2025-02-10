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
		"explore": {
			name:        "explore <location_area>",
			description: "View encounters in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch the specified pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
	}
}
