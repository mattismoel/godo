# Godo - The Simple CLI Todo List

This project aims to create a simple intuitive todo list for the terminal. It provides simple commands for manipulating one or multiple todos at a time with ease.


# Installation

To install this tool, one must have Golang installed on their system. After this, this project can be cloned to a desired location, and installed with the following steps:

```bash
$ go build -o "<bin_dir>/godo" main.go
```

The tool can now be used with the `godo` command.


Alternatively run the program with the Golang runtime directly:

```bash
$ go run main.go --help
```

# Usage

> Help for any command is provided with `godo --help` or `godo <command> --help`.

The program implements four different commands: `add`, `ls`, `toggle` and `rm`, for different todo manipulation:


To list the todos in the todo list, run the following command:
```bash
$ godo add "Grocery shopping" "Exercise" "Drink water" "Program in Go" # Add four different todo list entries.
```

To list the todos in the todo list run the following command:
```bash
$ godo ls # List all todos in todo list.
```
> The list command lists the ID of a given todo within `[#<id>]`.

Toggle todos:
```bash
$ godo toggle 1 3 2 5 # Toggle the 'done' state of todos with ID 1, 3, 2 and 5.
```

Remove todos from the todo list:
```bash
$ godo rm 1 3 2 5 # Remove todos with ID 1, 3, 2 and 5.
```
