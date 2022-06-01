## The idea is simple

You don't code for a certain amount of time, your code gets deleted.

Inspired by <a href="https://writeordie.com/">write-or-die</a>

<b> USE THIS AT YOUR OWN EXPENSE. </b> 

## Requirements

- Go (<a href="https://go.dev/doc/install">How to install?</a>)
- libasound2-dev (sudo apt install libasound2-dev)


## Installation

```
$ git clone https://github.com/Quimzy/code-or-die

$ cd code-or-die/

$ go build .

$ ./cod

```

## Usage

```
Usage of ./cod:
  -d string
        set direcotry (default <current dir>)
  -t int
        set interval time in minutes > 1 minute (default 30)
  -git (optional)
        set git mode

```
## Features

- [ ] Modes
- [X] Git support
- [X] Reminder
- [X] Windows Support

**I love beautiful pull requests and contributions. You got any ideas for new features? Implement it and pull request or add it as a feature <a href="https://github.com/Quimzy/code-or-die#features">here</a>**