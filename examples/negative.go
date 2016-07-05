package main

import (
	"fmt"

	"github.com/Lanzafame/filters"
)

// START

func negative(n int) int {
	return -n
}

func main() {
	naturals := make(chan int)
	negatives := make(chan int)

	go counter(naturals)
	go filters.IntFilter(negative, negatives, naturals)
	printer(negatives)
}

// END

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
