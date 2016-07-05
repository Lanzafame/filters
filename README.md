# Filters

Simple go channel pipeline filter operations package

### Example

The filter function:
[embedmd]:# (examples/negative.go /func negative/ /\}/)
```go
func negative(n int) int {
	return -n
}
```

Using the filter function:
[embedmd]:# (examples/negative.go /func main()/ /\}/)
```go
func main() {
	naturals := make(chan int)
	negatives := make(chan int)

	go counter(naturals)
	go filters.IntFilter(negative, negatives, naturals)
	printer(negatives)
}
```

