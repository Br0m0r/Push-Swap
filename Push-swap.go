package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Print("Not enough arguments.Please provide a stack example: go run Push-swap.go \" 3 2 1 3 2 \" ")
		return
	}

	input := os.Args[1]
	args := strings.Fields(input)

	// Parse input and validate
	stackA, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing the args")
		return
	}

	if isSorted(stackA) {
		fmt.Print("The stack is already sorted")
		return // Already sorted, no instructions needed
	}

	stackB := []int{}

	instructions := sortFive(stackA, stackB)

	// Output the instructions
	for _, instr := range instructions {
		fmt.Println(instr)
	}
	fmt.Println(len(instructions))
}

func shortTwo(stack []int) []string {
	var instructions []string

	if stack[0] > stack[1] {
		stack[0], stack[1] = stack[1], stack[0]
		instructions = append(instructions, "sa")
	}
	return instructions
}

func shortThree(stack []int) []string {
	var instructions []string
	MinIndex := findMinIndex(stack)

	if MinIndex == 0 && stack[1] > stack[2] {
		instructions = append(instructions, "sa", "ra")
		stack[0], stack[1] = stack[1], stack[0] // sa
		stack = rotateUp(stack)                 // ra
		return instructions
	}

	if MinIndex == 1 && stack[0] < stack[2] {
		instructions = append(instructions, "sa")
		stack[0], stack[1] = stack[1], stack[0] // sa
		return instructions
	}

	if MinIndex == 1 && stack[0] > stack[2] {
		instructions = append(instructions, "ra")
		stack = rotateUp(stack) // ra
		return instructions
	}
	if MinIndex == 2 && stack[0] < stack[1] {
		instructions = append(instructions, "rra")
		stack = rotateDown(stack) // rra
		return instructions
	}

	if MinIndex == 2 && stack[0] > stack[1] {
		instructions = append(instructions, "sa", "rra")
		stack[0], stack[1] = stack[1], stack[0] // sa
		stack = rotateDown(stack)               // rra
		return instructions
	}

	return instructions
}

// SortFive sorts up to five elements using stacks
func sortFive(stackA, stackB []int) []string {
	var instructions []string

	if stackA[0] > stackA[1] {
		// Find the two smallest elements in stackA
		smallest, secondSmallest := findTwoSmallest(stackA)

		if stackA[0] == secondSmallest && stackA[1] == smallest {
			instructions = append(instructions, "sa")
			stackA[0], stackA[1] = stackA[1], stackA[0] // Swap the first two elementsgo
		}
	}

	// Step 2: Move the smallest elements to stackB
	for len(stackA) > 3 {
		minIndex := findMinIndex(stackA)

		// Efficiently rotate in the correct direction
		for minIndex != 0 {
			if minIndex <= len(stackA)/2 {
				instructions = append(instructions, "ra")
				stackA = rotateUp(stackA)
			} else {
				instructions = append(instructions, "rra")
				stackA = rotateDown(stackA)
			}
			minIndex = findMinIndex(stackA)
		}

		// Push the smallest element to stackB
		instructions = append(instructions, "pb")
		stackB = append([]int{stackA[0]}, stackB...)
		stackA = stackA[1:]
	}

	// Step 3: Sort the remaining 3 elements in stackA
	if len(stackA) == 2 {
		instructions = append(instructions, shortTwo(stackA)...)
	} else {
		instructions = append(instructions, shortThree(stackA)...)
	}

	// Step 4: Move elements back from stackB to stackA
	for len(stackB) > 0 {
		instructions = append(instructions, "pa")
		stackA = append([]int{stackB[0]}, stackA...)
		stackB = stackB[1:]
	}

	return instructions
}

// Helper function to find the two smallest elements in the stack
func findTwoSmallest(stack []int) (int, int) {
	if len(stack) < 2 {
		return 0, 0 // Return default values if the stack is too small (less than 2 elements)
	}

	// Initialize the smallest and second smallest values
	var smallest, secondSmallest int
	if stack[0] < stack[1] {
		smallest = stack[0]
		secondSmallest = stack[1]
	} else {
		smallest = stack[1]
		secondSmallest = stack[0]
	}

	// Loop through the remaining elements in the stack to find the two smallest
	for i := 2; i < len(stack); i++ {
		if stack[i] < smallest {
			secondSmallest = smallest
			smallest = stack[i]
		} else if stack[i] < secondSmallest {
			secondSmallest = stack[i]
		}
	}

	return smallest, secondSmallest
}

// Parse and validate input
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

// Check if a stack is sorted
func isSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] > stack[i] {
			return false
		}
	}
	return true
}

// Rotate stack upwards
func rotateUp(stack []int) []int {
	return append(stack[1:], stack[0])
}

// Rotate stack downwards
func rotateDown(stack []int) []int {
	return append([]int{stack[len(stack)-1]}, stack[:len(stack)-1]...)
}

// Find the index of the minimum element
func findMinIndex(stack []int) int {
	minIndex := 0
	for i, val := range stack {
		if val < stack[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
