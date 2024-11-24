# Swissknife

🇨🇭🔪

Because everything in a single binary is noice.

This is a WIP 🚧.

## Build

```bash
# native
go build .

# linux/amd64
GOOS=linux GOARCH=amd64 go build .
```

## Usage

```bash
❯ swissknife
Various utils

Usage:
  swissknife [command]

Available Commands:
  completion         Generate the autocompletion script for the specified shell
  generate           Generate stuff
  get                Obtain information
  help               Help about any command
  list               List devices and volumes
  shouldideploytoday Should I deploy today?
  speedtest          Speedtest
  stopwatch          Create a stopwatch
  timer              Create a timer

Flags:
  -h, --help      help for swissknife
  -t, --toggle    Help message for toggle
  -v, --version   version for swissknife

Use "swissknife [command] --help" for more information about a command.
```
