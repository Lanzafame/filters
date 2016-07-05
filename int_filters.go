package filters

func IntFilter(f func(int) int, out chan<- int, in <-chan int) {
	defer close(out)
	for v := range in {
		out <- f(v)
	}
}
