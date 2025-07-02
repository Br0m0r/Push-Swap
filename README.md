# Push-Swap

A non-comparative sorting algorithm implementation using two stacks and a set of instructions.

## Overview

This project consists of two Go programs:
- **push-swap**: Generates the minimum number of instructions to sort a stack of integers
- **checker**: Validates if a sequence of instructions correctly sorts the given stack

## Available Instructions

| Instruction | Description |
|-------------|-------------|
| `pa` | Push top element from stack B to stack A |
| `pb` | Push top element from stack A to stack B |
| `sa` | Swap first 2 elements of stack A |
| `sb` | Swap first 2 elements of stack B |
| `ss` | Execute `sa` and `sb` |
| `ra` | Rotate stack A up (first becomes last) |
| `rb` | Rotate stack B up |
| `rr` | Execute `ra` and `rb` |
| `rra` | Reverse rotate stack A (last becomes first) |
| `rrb` | Reverse rotate stack B |
| `rrr` | Execute `rra` and `rrb` |

## Usage

### Build the programs
```bash
go build -o push-swap Push-swap.go
go build -o checker checker.go
```

### Running push-swap
```bash
./push-swap "2 1 3 6 5 8"
```
Output: List of instructions to sort the stack

### Running checker
```bash
echo -e "sa\nrra\npb" | ./checker "3 2 1 0"
```
Output: `OK` if sorted correctly, `KO` otherwise

### Combined usage
```bash
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
```

## Error Handling

Both programs display `Error` on stderr for:
- Non-integer arguments
- Duplicate numbers
- Invalid instructions (checker only)

## Requirements

- Written in Go using only standard library packages
- Minimize the number of sorting instructions
- Handle all edge cases and errors appropriately

## Goal

Sort stack A in ascending order with the minimum possible number of operations using both stacks and the available instruction set.
