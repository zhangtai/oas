# Overcomplicated Automation System

## Base

### reminders-cli

This tool is required to export reminders items as JSON and read in base api endpoint.

1. Clone [0xdevalias/reminders-cli](https://github.com/0xdevalias/reminders-cli/)
1. `cd reminders-cli && make build-release`
1. `sudo cp .build/release/reminders /usr/local/bin/reminders`

## Launcher

Open Alfred -> Advanced -> Set preferences folder, and set to `./launcher`

## Deck

- Clone submodule: `git submodule update --init --recursive`
- Link the plugin: `ln -sf ${PWD}/deck/ai.zhangt.oas.sdPlugin ~/Library/Application\ Support/com.elgato.StreamDeck/Plugins/ai.zhangt.oas.sdPlugin`

## Keyboard Shortcuts

Check [./keyboard_shortcuts.md](keyboard_shortcuts.md)
