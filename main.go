package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scn  = bufio.NewReader(os.Stdin)
	prnt = bufio.NewWriter(os.Stdout)
)

func solve() {
}

func main() {
	defer prnt.Flush()
}

func min(a ...int) int {
	mn := a[0]
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] < mn {
			mn = a[i]
		}
	}
	return mn
}

func max(a ...int) int {
	mx := a[0]
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] > mx {
			mx = a[i]
		}
	}
	return mx
}

func queue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []int) []int {
	return queue[1:]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
