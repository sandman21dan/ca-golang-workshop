# ca-golang-workshop

This is a repository that contains basic scaffolding to be able to follow along the
Go workshop for the NASA event

Welcome to NASA's hackathon/workshop

We'll be challenging you to solve some problems using GoLang

## Table of contents

- [ca-golang-workshop](#ca-golang-workshop)
  * [Pre-requisites](#pre-requisites)
    + [Install Go](#install-go)
      - [Install goenv](#install-goenv)
      - [Install a Go version](#install-a-go-version)
    + [The Docker way](#the-docker-way)
  * [Exercises](#exercises)
    + [Track 1](#track-1)
      - [Roman](#roman)
      - [Bowling scores](#bowling-scores)
      - [Password generator](#password-generator)

## Pre-requisites

Everything can be run using Docker, but you'll likely have a better
development experience if you have go installed locally

### Install Go

#### Install goenv

Download goenv
```sh
git clone https://github.com/syndbg/goenv.git ~/.goenv
```

Add goenv to PATH
```sh
echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bashrc
echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(goenv init -)"' >> ~/.bashrc
```

To activate goenv either reload your terminal or run:

```sh
source ~/.bashrc
```

**Note:** If you're using `zsh` instead of `bash` replace `~/.bashrc` for `~/.zshrc` on all previous commands

#### Install a Go version

```sh
goenv install 1.17.7
```

Make that version the default

```sh
goenv global 1.17.7
```

Test that the go command is working

```sh
go version
```

### The Docker way

Run the following command on the root directory of this project

```sh
docker compose up -d
```

To run any of the go commands below just prepend the following

```sh
docker exec go-runner <COMMAND>
```

For example

```sh
docker exec go-runner go test -run ^TestRomanNumerals$ ./...
```

## Exercises

### Track 1

#### Roman

[Start here](./app/pkg/track1/roman/roman.go)
[Test cases](./app/pkg/track1/roman/roman_test.go)

Create a function that converts from roman number notation to decimal

Function should work as follows

```go
roman.ToDecimal("XXIV") // 14
```

More info on Roman numerals [here](https://www.dictionary.com/e/roman-numerals)

Run test cases
```sh
go test -run ^TestRomanNumerals$ ./...
```

#### Bowling scores

[Start here](./app/pkg/track1/bowling/bowling.go)
[Test cases](./app/pkg/track1/bowling/bowling_test.go)

Create a function correctly calculates a bowling score sheet

No validation is expected, assume the right amount of bowling cells are provided
and no illegal combinations are provided

Function should work as follows

```go
bowling.CalculateBowlingScore("X X X X X X X X X X X X") // 300
```

More info on bowling scoring [here](https://codingdojo.org/kata/Bowling/)

Run test cases
```sh
go test -run ^TestCalculateBowlingScore$ ./...
```

#### Password generator

[Start here](./app/pkg/track1/passwordgenerator/password_generator.go)
[Test cases](./app/pkg/track1/passwordgenerator/password_generator_test.go)

Create a function that randomly generates passwords

The function takes an int as first argument for length and
an option `struct` as a second argument that controls the options
of the password to be generated

At least one character of each type passed in the options is expected to be
in the generated password

Don't worry about producing cryptographically secure randomness any kind
of random function suffices for this exercise

Function should work as follows

```go
passwordgenerator.GeneratePassword(10, GeneratePasswordOptions{ Digits: true, Letters: true, Uppercase: true, SpecialCharacters: true}) // f$dZ4Ui#bl
```

Run test cases
```sh
go test -run ^TestGeneratePassword$ ./...
```
