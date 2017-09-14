package cooperative

import "fmt"

// diff from "before 1.5" and "after 1.5"
func BlockByFor() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println("hello - ", index)
		}(i)
	}

	for {}
}
