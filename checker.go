//package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return // No arguments, no output
	}

	input := os.Args[1]
	args := strings.Fields(input)

	stackA, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	stackB := []int{}

	// Read instructions from standard input
	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// Apply instructions
	for _, instr := range instructions {
		switch instr {
		case "pa":
			if len(stackB) > 0 {
				stackA = append([]int{stackB[0]}, stackA...)
				stackB = stackB[1:]
			}
		case "pb":
			if len(stackA) > 0 {
				stackB = append([]int{stackA[0]}, stackB...)
				stackA = stackA[1:]
			}
		case "sa":
			if len(stackA) > 1 {
				stackA[0], stackA[1] = stackA[1], stackA[0]
			}
		case "sb":
			if len(stackB) > 1 {
				stackB[0], stackB[1] = stackB[1], stackB[0]
			}
		case "ss":
			if len(stackA) > 1 {
				stackA[0], stackA[1] = stackA[1], stackA[0]
			}
			if len(stackB) > 1 {
				stackB[0], stackB[1] = stackB[1], stackB[0]
			}
		case "ra":
			stackA = rotateUp(stackA)
		case "rb":
			stackB = rotateUp(stackB)
		case "rr":
			stackA = rotateUp(stackA)
			stackB = rotateUp(stackB)
		case "rra":
			stackA = rotateDown(stackA)
		case "rrb":
			stackB = rotateDown(stackB)
		case "rrr":
			stackA = rotateDown(stackA)
			stackB = rotateDown(stackB)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	// Check if stackA is sorted and stackB is empty
	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// Helper functions (parseArgs, isSorted, rotateUp, rotateDown) as defined earlier
func parseArgs(args []string) ([]int, error) {
	stack := make([]int, len(args))
	seen := map[int]bool{}

	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil || seen[num] {
			return nil, fmt.Errorf("invalid input")
		}
		stack[i] = num
		seen[num] = true
	}
	return stack, nil
}

func isSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] > stack[i] {
			return false
		}
	}
	return true
}

func rotateUp(stack []int) []int {
	return append(stack[1:], stack[0])
}

func rotateDown(stack []int) []int {
	return append([]int{stack[len(stack)-1]}, stack[:len(stack)-1]...)
}
