package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing to recover. " +
				"Please try uncomment errors " +
				"below.")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
}

func main() {
	tryRecover()
}
