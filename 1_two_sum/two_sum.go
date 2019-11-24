package main

import (
	"fmt"
	"log"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	fmt.Println("Successfully ran")
}

func Run() error {
	return nil
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)-1)
	for j, a := range nums {
		if i, ok := m[target-a]; ok {
			return []int{i, j}
		}
		m[a] = j
	}
	return []int{}
}
