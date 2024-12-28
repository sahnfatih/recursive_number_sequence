# Recursive Number Sequence Generator

A Go application that generates a specific number sequence using recursion.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Testing](#testing)

## Features
- Generates a sequence of numbers based on recursive division by 2
- Command-line interface
- Error handling for invalid inputs
- Comprehensive test coverage

## Project Structure
recursive_number_sequence/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   └── sequence/
│       ├── sequence.go    # Core sequence logic
│       └── sequence_test.go # Unit tests
├── go.mod              # Go module file
└── README.md          # Documentation

## Requirements
- Go 1.20 or higher
- Git

## Installation
1. Clone the repository:
git clone https://github.com/sahnfatih/recursive_number_sequence.git

2. Navigate to project directory:
cd recursive_number_sequence

3. Build the project:
go build ./cmd/main.go
## Algorithm Description

The program implements a recursive algorithm that:
1. Takes an integer input n
2. Recursively divides the number by 2 until reaching 1
3. During the recursive callbacks, builds a sequence of numbers
4. Only includes numbers greater than 1 in the result

Example for input 9:
## Usage
Run the program with a number:
go run cmd/main.go 9

Example output:
Input: 9
Output:
2
4
9

## Development
### Code Structure
- All function and variable names follow Go naming conventions
- Comments are in English and follow Go best practices
- Core logic is separated into the internal package

## Testing
Run all tests:
go test ./...

Run tests with coverage:
go test ./... -cover