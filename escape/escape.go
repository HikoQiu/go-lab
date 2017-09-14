package main

func foo() *int {
	var x int
	return &x
}

func bar() int {
	b := new(int)
	*b = 1
	return *b
}

// go run -gcflags '-m -l' src/sharing/escape/escape.go
func main()  {

	foo()
	bar()
}
