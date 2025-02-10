<p align="center">
    <img src="https://github.com/user-attachments/assets/dd3fb513-1e3a-4580-9778-aae4759c87e3" align="center" width="30%">
</p>
<p align="center"><h1 align="center">POKEDEXCLI</h1></p>
<p align="center">
    <em><code>❯ A Command Line Interface for Exploring the Pokémon World</code></em>
</p>
<p align="center">
    <img src="https://img.shields.io/github/license/abhinaenae/pokedexcli?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
    <img src="https://img.shields.io/github/last-commit/abhinaenae/pokedexcli?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
    <img src="https://img.shields.io/github/languages/top/abhinaenae/pokedexcli?style=default&color=0080ff" alt="repo-top-language">
    <img src="https://img.shields.io/github/languages/count/abhinaenae/pokedexcli?style=default&color=0080ff" alt="repo-language-count">
</p>

---

## Overview

PokedexCLI is a simple yet powerful command-line tool that allows users to interact with Pokémon data. Using the PokéAPI, this tool lets you explore various Pokémon species, their abilities, types, and habitats directly from your terminal.

---

## Features

- Search for Pokémon by name
- View detailed information including abilities, types, and stats
- Explore different locations where Pokémon can be found
- Catch and store Pokémon in your personal Pokédex
- View a map of recently explored locations
- Easy-to-use command-line interface

---

## Project Structure

```sh
└── pokedexcli/
    ├── command_catch.go        # Handles catching Pokémon
    ├── command_exit.go         # Handles exiting the CLI
    ├── command_explore.go      # Handles exploring locations
    ├── command_help.go         # Displays available commands
    ├── command_inspect.go      # Displays details about a Pokémon
    ├── command_map.go          # Displays a map of visited locations
    ├── command_view.go         # Shows caught Pokémon
    ├── commands.go             # Defines CLI commands
    ├── go.mod                  # Go module dependencies
    ├── helpers/
    │   ├── cleaner.go          # Utility functions for cleaning user input
    │   └── cleaner_test.go     # Unit tests for cleaner functions
    ├── init.go                 # Initialization logic
    ├── internal/
    │   ├── pokeapi/            # Handles API requests to PokéAPI
    │   ├── pokecache/          # Implements caching for API requests
    ├── main.go                 # Entry point of the CLI application
    └── repl.go                 # Handles the command-line REPL interface
```

---

## Getting Started

### Prerequisites

Before getting started with pokedexcli, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go 1.23
- **Package Manager:** Go modules

### Installation

Install pokedexcli using one of the following methods:

**Build from source:**

1. Clone the pokedexcli repository:
```sh
❯ git clone https://github.com/abhinaenae/pokedexcli
```

2. Navigate to the project directory:
```sh
❯ cd pokedexcli
```

3. Install the project dependencies:

**Using `go modules`**
```sh
❯ go build
```

### Usage

Run pokedexcli using the following command:
```sh
❯ ./pokedexcli
```

Once inside the CLI, use the following commands to explore Pokémon data:

- `explore <location>` – View Pokémon in a specific location
- `inspect <pokemon>` – Get details about a Pokémon
- `catch <pokemon>` – Attempt to catch a Pokémon
- `map` – Display visited locations
- `view` – Show caught Pokémon
- `help` – List available commands
- `exit` – Quit the CLI

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

---

## License

This project is licensed under the MIT License.

---

## Acknowledgments

- [PokéAPI](https://pokeapi.co/) for providing the Pokémon data
- Inspired by various CLI-based projects

