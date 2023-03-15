# bubble-tictactoe

A simple Tic Tac Toe played on the terminal using library [Bubbletea](https://github.com/charmbracelet/bubbletea) from [charm.sh](https://charm.sh).

## Install

```sh
go install github.com/christopher-kleine/bubble-tictactoe/cmd/bubble-tictactoe@latest
```

## Usage

You can press one of the following keys:

- h = Show the help
- t = Toggle through the board designs
- Up, Down, Left, Right = Move the cursor
- Enter = Place your piece on the board
- q, Ctrl+C = Quit the game
- n = Start a new game
- m = Let the AI make a move

## Purpose

This simple game is meant as part of a small game collection that can be played over SSH. For this, the repo [buuble-games](https://github.com/christopher-kleine.de/bubble-games) exist.

However, the game can be played standalone. See the [Install](#install) section for this.
